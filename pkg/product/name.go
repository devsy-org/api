package product

import (
	"fmt"
	"os"
	"sync"

	"github.com/devsy-org/admin-apis/pkg/licenseapi"
	"k8s.io/klog/v2"
)

// Product is the global variable to be set at build time
var (
	productName string = string(licenseapi.DevsyOrg)
	once        sync.Once
)

func loadProductVar() {
	productEnv := os.Getenv("PRODUCT")
	switch {
	case productEnv == string(licenseapi.DevsyPro):
		productName = string(licenseapi.DevsyPro)
	case productEnv == string(licenseapi.DevsyOrg):
		productName = string(licenseapi.DevsyOrg)
	case productEnv != "":
		klog.TODO().
			Error(fmt.Errorf("unrecognized product %s", productEnv), "error parsing product", "product", productEnv)
	}
}

func Name() licenseapi.ProductName {
	once.Do(loadProductVar)
	return licenseapi.ProductName(productName)
}

// DisplayName returns the display name of the product.
func DisplayName() string {
	switch Name() {
	case licenseapi.DevsyPro:
		return "Devsy Pro"
	case licenseapi.DevsyOrg:
		return "Devsy"
	}

	return "Devsy"
}
