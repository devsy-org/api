package product

import (
	"fmt"
	"os"
	"sync"

	"github.com/skevetter/admin-apis/pkg/licenseapi"
	"k8s.io/klog/v2"
)

// Product is the global variable to be set at build time
var (
	productName string = string(licenseapi.Devsy)
	once        sync.Once
)

func loadProductVar() {
	productEnv := os.Getenv("PRODUCT")
	if productEnv == string(licenseapi.DevPodPro) {
		productName = string(licenseapi.DevPodPro)
	} else if productEnv == string(licenseapi.DevsyPro) {
		productName = string(licenseapi.DevsyPro)
	} else if productEnv == string(licenseapi.Devsy) {
		productName = string(licenseapi.Devsy)
	} else if productEnv != "" {
		klog.TODO().
			Error(fmt.Errorf("unrecognized product %s", productEnv), "error parsing product", "product", productEnv)
	}
}

func Name() licenseapi.ProductName {
	once.Do(loadProductVar)
	return licenseapi.ProductName(productName)
}

// Name returns the name of the product
func DisplayName() string {
	devsyDisplayName := "Devsy"

	switch Name() {
	case licenseapi.DevPodPro:
		return "DevPod Pro"
	case licenseapi.DevsyPro:
		return "vCluster Platform"
	case licenseapi.Devsy:
	}

	return devsyDisplayName
}
