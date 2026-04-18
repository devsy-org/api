package product

import (
	"fmt"
	"strings"

	"github.com/devsy-org/admin-apis/pkg/licenseapi"
)

// Replace replaces the product name in the given usage string
// based on the current product.Product().
//
// It replaces "devsy" with the specific product name:
//   - "devsy platform" for product.DevsyPro
//   - No replacement for product.DevsyOrg
//
// Parameters:
//   - content: The string to update
//
// Returns:
//   - The updated string with product name replaced if needed.
func Replace(content string) string {
	switch Name() {
	case licenseapi.DevsyPro:
		content = strings.Replace(content, "devsy.sh", "devsy.pro", -1)
		content = strings.Replace(content, "devsy.host", "devsy.host", -1)

		content = strings.Replace(content, "devsy", "devsy platform", -1)
		content = strings.Replace(content, "Devsy", "vCluster Platform", -1)
	case licenseapi.DevsyOrg:
	}

	return content
}

// ReplaceWithHeader replaces the "devsy" product name in the given
// usage string with the specific product name based on product.Product().
// It also adds a header with padding around the product name and usage.
//
// The product name replacements are:
//
//   - "devsy platform" for product.DevsyPro
//   - No replacement for product.DevsyOrg
//
// Parameters:
//   - use: The usage string
//   - content: The content string to run product name replacement on
//
// Returns:
//   - The content string with product name replaced and header added
func ReplaceWithHeader(use, content string) string {
	maxChar := 56

	productName := licenseapi.DevsyOrg

	switch Name() {
	case licenseapi.DevsyPro:
		productName = "devsy platform"
	case licenseapi.DevsyOrg:
	}

	paddingSize := (maxChar - 2 - len(productName) - len(use)) / 2

	separator := strings.Repeat("#", paddingSize*2+len(productName)+len(use)+2+1)
	padding := strings.Repeat("#", paddingSize)

	return fmt.Sprintf(`%s
%s %s %s %s
%s
%s
`, separator, padding, productName, use, padding, separator, Replace(content))
}
