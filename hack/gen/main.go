package main

import (
	"os"
	"runtime"

	"github.com/devsy-org/apiserver/pkg/generate"
	"k8s.io/klog/v2"
)

func main() {
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	g := generate.Gen{}
	if err := g.Execute("zz_generated.api.register.go", "github.com/devsy-org/api/pkg/apis/..."); err != nil {
		klog.Fatalf("Error: %v", err)
	}
	klog.V(2).Info("Completed successfully.")
}
