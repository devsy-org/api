// hack/generate.go orchestrates all code generation for the devsy-org/api module.
//
// Usage:
//
//	go run ./hack/generate.go              # Run all generators
//	go run ./hack/generate.go register     # Run only API register generation
//	go run ./hack/generate.go --help       # Show all commands
package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/devsy-org/apiserver/pkg/generate"
	"github.com/urfave/cli/v3"
	"k8s.io/klog/v2"
)

const module = "github.com/devsy-org/api"

var (
	boilerplate = filepath.Join(repoRoot(), "hack", "boilerplate.go.txt")

	apiPackages = []string{
		module + "/pkg/apis/audit/v1",
		module + "/pkg/apis/management",
		module + "/pkg/apis/management/v1",
		module + "/pkg/apis/storage/v1",
		module + "/pkg/apis/ui",
		module + "/pkg/apis/ui/v1",
		module + "/pkg/apis/virtualcluster",
		module + "/pkg/apis/virtualcluster/v1",
	}

	clientInputDirs = "management/v1,storage/v1,virtualcluster/v1"

	// versionedPackages are the versioned API packages (no hub packages).
	// Used by lister-gen and informer-gen which require versioned input only.
	versionedPackages = []string{
		module + "/pkg/apis/audit/v1",
		module + "/pkg/apis/management/v1",
		module + "/pkg/apis/storage/v1",
		module + "/pkg/apis/ui/v1",
		module + "/pkg/apis/virtualcluster/v1",
	}

	conversionPackages = []string{
		module + "/pkg/apis/management/v1",
		module + "/pkg/apis/virtualcluster/v1",
	}

	openapiExtra = []string{
		"k8s.io/apimachinery/pkg/apis/meta/v1",
		"k8s.io/apimachinery/pkg/api/resource",
		"k8s.io/apimachinery/pkg/version",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/util/intstr",
		"k8s.io/api/core/v1",
		"k8s.io/api/rbac/v1",
		"k8s.io/api/apps/v1",
		"k8s.io/api/networking/v1",
		"k8s.io/api/storage/v1",
		"k8s.io/api/batch/v1",
	}
)

func repoRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filename))
}

func main() {
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	os.Chdir(repoRoot())

	app := &cli.Command{
		Name:  "generate",
		Usage: "Code generation for devsy-org/api",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return runAll()
		},
		Commands: []*cli.Command{
			{
				Name:  "all",
				Usage: "Run all generators (default)",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return runAll()
				},
			},
			{
				Name:  "install",
				Usage: "Install code-generator tool binaries",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					installTools()
					return nil
				},
			},
			{
				Name:  "register",
				Usage: "Generate API registration (zz_generated.api.register.go)",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					runRegister()
					return nil
				},
			},
			{
				Name:  "deepcopy",
				Usage: "Generate DeepCopy methods (zz_generated.deepcopy.go)",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					runGenerator("deepcopy-gen", "--go-header-file", boilerplate, "--output-file", "zz_generated.deepcopy.go")
					return nil
				},
			},
			{
				Name:  "defaults",
				Usage: "Generate Defaulter functions (zz_generated.defaults.go)",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					runGenerator("defaulter-gen", "--go-header-file", boilerplate, "--output-file", "zz_generated.defaults.go")
					return nil
				},
			},
			{
				Name:  "conversion",
				Usage: "Generate Conversion functions (zz_generated.conversion.go)",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					runConversion()
					return nil
				},
			},
			{
				Name:  "openapi",
				Usage: "Generate OpenAPI definitions (zz_generated.openapi.go)",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					runOpenAPI()
					return nil
				},
			},
			{
				Name:  "clients",
				Usage: "Generate clientset, listers, and informers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					runClients()
					return nil
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		klog.Fatalf("Error: %v", err)
	}
}

func runAll() error {
	installTools()
	runRegister()
	runGenerator("deepcopy-gen", "--go-header-file", boilerplate, "--output-file", "zz_generated.deepcopy.go")
	runGenerator("defaulter-gen", "--go-header-file", boilerplate, "--output-file", "zz_generated.defaults.go")
	runConversion()
	runOpenAPI()
	runClients()
	fmt.Println("==> Done.")
	return nil
}

// installTools installs the k8s.io/code-generator binaries.
func installTools() {
	fmt.Println("==> Installing code-generator tools...")
	tools := []string{
		"k8s.io/code-generator/cmd/deepcopy-gen",
		"k8s.io/code-generator/cmd/defaulter-gen",
		"k8s.io/code-generator/cmd/conversion-gen",
		"k8s.io/code-generator/cmd/client-gen",
		"k8s.io/code-generator/cmd/lister-gen",
		"k8s.io/code-generator/cmd/informer-gen",
		"k8s.io/kube-openapi/cmd/openapi-gen",
	}
	for _, tool := range tools {
		run("go", "install", tool)
	}
}

// runRegister runs the custom API register generator from apiserver/pkg/generate
// with a GroupConverter that deduplicates import aliases.
//
// The upstream generator has two known bugs that are patched here:
//
//  1. Import alias collisions: it joins the last two path segments to form aliases
//     (e.g. "storage"+"v1" → "storagev1"), producing duplicates when two modules
//     share that suffix. Fixed via a GroupConverter hook.
//
//  2. NewRESTFunc signature: the template emits `func() rest.Storage` but subresource
//     closures call `RESTFunc(Factory)`. The correct signature is
//     `func(managerfactory.SharedManagerFactory) rest.Storage`. Fixed via post-processing.
func runRegister() {
	fmt.Println("==> Generating API register...")

	g := generate.Gen{
		GroupConverter: fixImportAliases,
	}
	if err := g.Execute("zz_generated.api.register.go", module+"/pkg/apis/..."); err != nil {
		klog.Fatalf("register generation failed: %v", err)
	}

	patchRegisterFiles()
}

// patchRegisterFiles fixes the NewRESTFunc signature and adds the managerfactory
// import to generated register files in packages that declare a Factory variable.
func patchRegisterFiles() {
	targets := []string{
		filepath.Join("pkg", "apis", "management", "zz_generated.api.register.go"),
		filepath.Join("pkg", "apis", "virtualcluster", "zz_generated.api.register.go"),
	}

	mfImport := fmt.Sprintf(`"%s/pkg/managerfactory"`, module)

	for _, path := range targets {
		data, err := os.ReadFile(path)
		if err != nil {
			klog.Warningf("skipping patch for %s: %v", path, err)
			continue
		}
		content := string(data)

		// 1. Fix NewRESTFunc signature.
		content = strings.Replace(content,
			"type NewRESTFunc func() rest.Storage",
			"type NewRESTFunc func(factory managerfactory.SharedManagerFactory) rest.Storage",
			1)

		// 2. Fix top-level REST closures missing the Factory arg.
		restFuncCallRe := regexp.MustCompile(`return (New\w+RESTFunc)\(\)`)
		content = restFuncCallRe.ReplaceAllString(content, "return ${1}(Factory)")

		// 3. Add managerfactory import if not already present.
		if !strings.Contains(content, mfImport) {
			content = strings.Replace(content,
				"import (",
				"import (\n\t"+mfImport,
				1)
		}

		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			klog.Fatalf("failed to write patched %s: %v", path, err)
		}

		// Run goimports to fix formatting after our edits.
		goimportsBin := "goimports"
		if p := filepath.Join(gobin(), "goimports"); fileExists(p) {
			goimportsBin = p
		}
		fmtCmd := exec.Command(goimportsBin, "-w", path)
		if out, err := fmtCmd.CombinedOutput(); err != nil {
			fmtCmd = exec.Command("gofmt", "-w", path)
			if out2, err2 := fmtCmd.CombinedOutput(); err2 != nil {
				klog.Warningf("gofmt %s: %s %v", path, string(append(out, out2...)), err2)
			}
		}
	}
}

// fixImportAliases walks every struct field in the API group and rewrites
// any duplicate import aliases so they are unique.
func fixImportAliases(apigroup *generate.APIGroup) {
	seen := map[string]string{}
	importRe := regexp.MustCompile(`^(\w+)\s+"(.+)"$`)

	for _, s := range apigroup.Structs {
		for _, f := range s.Fields {
			if f.UnversionedImport == "" {
				continue
			}

			m := importRe.FindStringSubmatch(f.UnversionedImport)
			if m == nil {
				continue
			}
			alias, pkgPath := m[1], m[2]

			if first, exists := seen[alias]; exists && first != pkgPath {
				newAlias := disambiguateAlias(pkgPath, alias)
				f.UnversionedImport = fmt.Sprintf(`%s "%s"`, newAlias, pkgPath)
				f.UnversionedType = strings.Replace(f.UnversionedType, alias+".", newAlias+".", 1)
				seen[newAlias] = pkgPath
			} else {
				seen[alias] = pkgPath
			}
		}
	}
}

// disambiguateAlias creates a unique import alias by incorporating the module
// name from the package path. For "github.com/devsy-org/agentapi/pkg/apis/devsy/storage/v1"
// with base alias "storagev1", this produces "agentstoragev1".
func disambiguateAlias(pkgPath, baseAlias string) string {
	parts := strings.Split(pkgPath, "/")

	moduleName := ""
	if len(parts) >= 3 {
		moduleName = parts[2]
	}

	prefix := moduleName
	prefix = strings.ReplaceAll(prefix, "-", "")
	prefix = strings.TrimSuffix(prefix, "apis")
	prefix = strings.TrimSuffix(prefix, "api")

	if prefix == "" || prefix == baseAlias {
		prefix = strings.ReplaceAll(moduleName, "-", "")
	}

	return prefix + baseAlias
}

func runGenerator(tool string, baseArgs ...string) {
	fmt.Printf("==> Generating %s...\n", tool)
	args := append(baseArgs, apiPackages...)
	run(tool, args...)
}

func runConversion() {
	fmt.Println("==> Generating conversion...")
	args := []string{
		"--go-header-file", boilerplate,
		"--output-file", "zz_generated.conversion.go",
	}
	args = append(args, conversionPackages...)
	run("conversion-gen", args...)
}

func runOpenAPI() {
	fmt.Println("==> Generating openapi...")
	allPkgs := append(append([]string{}, apiPackages...), openapiExtra...)
	args := []string{
		"--go-header-file", boilerplate,
		"--output-pkg", module + "/pkg/openapi",
		"--output-file", "zz_generated.openapi.go",
		"--output-dir", "pkg/openapi",
		"--report-filename", "/dev/null",
	}
	args = append(args, allPkgs...)
	run("openapi-gen", args...)
}

func runClients() {
	fmt.Println("==> Generating clientset...")
	run("client-gen",
		"--go-header-file", boilerplate,
		"--input-base", module+"/pkg/apis",
		"--input", clientInputDirs,
		"--output-pkg", module+"/pkg/clientset",
		"--clientset-name", "versioned",
		"--output-dir", "pkg/clientset",
	)

	fmt.Println("==> Generating listers...")
	args := []string{
		"--go-header-file", boilerplate,
		"--output-pkg", module + "/pkg/listers",
		"--output-dir", "pkg/listers",
	}
	args = append(args, versionedPackages...)
	run("lister-gen", args...)

	fmt.Println("==> Generating informers...")
	args = []string{
		"--go-header-file", boilerplate,
		"--output-pkg", module + "/pkg/informers",
		"--output-dir", "pkg/informers",
		"--versioned-clientset-package", module + "/pkg/clientset/versioned",
		"--listers-package", module + "/pkg/listers",
	}
	args = append(args, versionedPackages...)
	run("informer-gen", args...)
}

// gobin returns the GOBIN directory (where `go install` places binaries).
func gobin() string {
	if b := os.Getenv("GOBIN"); b != "" {
		return b
	}
	out, _ := exec.Command("go", "env", "GOPATH").Output()
	return filepath.Join(strings.TrimSpace(string(out)), "bin")
}

// run executes a command, forwarding stdout/stderr. Exits on failure.
// It resolves the executable from GOBIN first, then falls back to PATH.
func run(name string, args ...string) {
	bin := name
	if p := filepath.Join(gobin(), name); fileExists(p) {
		bin = p
	}
	cmd := exec.Command(bin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		klog.Fatalf("%s failed: %v", name, err)
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
