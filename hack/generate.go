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
	"os"
	"os/exec"
	"path/filepath"
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
	if err := os.Chdir(repoRoot()); err != nil {
		klog.Fatalf("chdir failed: %v", err)
	}

	app := &cli.Command{
		Name:  "generate",
		Usage: "Code generation for devsy-org/api",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return runAll()
		},
		Commands: buildCommands(),
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		klog.Fatalf("Error: %v", err)
	}
}

func buildCommands() []*cli.Command {
	return []*cli.Command{
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
				runGenerator(
					"deepcopy-gen",
					"--go-header-file",
					boilerplate,
					"--output-file",
					"zz_generated.deepcopy.go",
				)
				return nil
			},
		},
		{
			Name:  "defaults",
			Usage: "Generate Defaulter functions (zz_generated.defaults.go)",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				runGenerator(
					"defaulter-gen",
					"--go-header-file",
					boilerplate,
					"--output-file",
					"zz_generated.defaults.go",
				)
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
	}
}

func runAll() error {
	installTools()
	runRegister()
	runGenerator(
		"deepcopy-gen",
		"--go-header-file",
		boilerplate,
		"--output-file",
		"zz_generated.deepcopy.go",
	)
	runGenerator(
		"defaulter-gen",
		"--go-header-file",
		boilerplate,
		"--output-file",
		"zz_generated.defaults.go",
	)
	runConversion()
	runOpenAPI()
	runClients()
	klog.Info("==> Done.")
	return nil
}

// installTools installs the k8s.io/code-generator binaries.
func installTools() {
	klog.Info("==> Installing code-generator tools...")
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

func runRegister() {
	klog.Info("==> Generating API register...")

	g := generate.Gen{}
	if err := g.Execute("zz_generated.api.register.go", module+"/pkg/apis/..."); err != nil {
		klog.Fatalf("register generation failed: %v", err)
	}
}

func runGenerator(tool string, baseArgs ...string) {
	klog.Infof("==> Generating %s...", tool)
	args := append([]string{}, baseArgs...)
	args = append(args, apiPackages...)
	run(tool, args...)
}

func runConversion() {
	klog.Info("==> Generating conversion...")
	args := []string{
		"--go-header-file", boilerplate,
		"--output-file", "zz_generated.conversion.go",
	}
	args = append(args, conversionPackages...)
	run("conversion-gen", args...)
}

func runOpenAPI() {
	klog.Info("==> Generating openapi...")
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
	klog.Info("==> Generating clientset...")
	run("client-gen",
		"--go-header-file", boilerplate,
		"--input-base", module+"/pkg/apis",
		"--input", clientInputDirs,
		"--output-pkg", module+"/pkg/clientset",
		"--clientset-name", "versioned",
		"--output-dir", "pkg/clientset",
	)

	klog.Info("==> Generating listers...")
	args := []string{
		"--go-header-file", boilerplate,
		"--output-pkg", module + "/pkg/listers",
		"--output-dir", "pkg/listers",
	}
	args = append(args, versionedPackages...)
	run("lister-gen", args...)

	klog.Info("==> Generating informers...")
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
	cmd := exec.Command(bin, args...) //nolint:gosec // G204: binary is resolved from GOBIN or PATH, not user input
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
