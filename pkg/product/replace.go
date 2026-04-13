package product

import (
	"fmt"
	"strings"

	"github.com/skevetter/admin-apis/pkg/licenseapi"
)

// Replace replaces the product name in the given usage string
// based on the current product.Product().
//
// It replaces "devsy" with the specific product name:
//   - "devpod pro" for product.DevPodPro
//   - "devsy platform" for product.DevsyPro
//   - No replacement for product.Devsy
//
// This handles case insensitive replaces like "devsy" -> "devpod pro".
//
// It also handles case sensitive replaces:
//   - "Devsy" -> "DevPod.Pro" for product.DevPodPro
//   - "Devsy" -> "vCluster Platform" for product.DevsyPro
//
// This allows customizing command usage text for different products.
//
// Parameters:
//   - content: The string to update
//
// Returns:
//   - The updated string with product name replaced if needed.
func Replace(content string) string {
	switch Name() {
	case licenseapi.DevPodPro:
		content = strings.Replace(content, "devsy.sh", "devpod.pro", -1)
		content = strings.Replace(content, "devsy.host", "devpod.host", -1)

		content = strings.Replace(content, "devsy", "devpod pro", -1)
		content = strings.Replace(content, "Devsy", "DevPod.Pro", -1)
	case licenseapi.DevsyPro:
		content = strings.Replace(content, "devsy.sh", "devsy.pro", -1)
		content = strings.Replace(content, "devsy.host", "devsy.host", -1)

		content = strings.Replace(content, "devsy", "devsy platform", -1)
		content = strings.Replace(content, "Devsy", "vCluster Platform", -1)
	case licenseapi.Devsy:
	}

	return content
}

// ReplaceWithHeader replaces the "devsy" product name in the given
// usage string with the specific product name based on product.Product().
// It also adds a header with padding around the product name and usage.
//
// The product name replacements are:
//
//   - "devpod pro" for product.DevPodPro
//   - "devsy platform" for product.DevsyPro
//   - No replacement for product.Devsy
//
// This handles case insensitive replaces like "devsy" -> "devpod pro".
//
// It also handles case sensitive replaces:
//   - "Devsy" -> "DevPod.Pro" for product.DevPodPro
//   - "Devsy" -> "vCluster Platform" for product.DevsyPro
//
// Parameters:
//   - use: The usage string
//   - content: The content string to run product name replacement on
//
// Returns:
//   - The content string with product name replaced and header added
func ReplaceWithHeader(use, content string) string {
	maxChar := 56

	productName := licenseapi.Devsy

	switch Name() {
	case licenseapi.DevPodPro:
		productName = "devpod pro"
	case licenseapi.DevsyPro:
		productName = "devsy platform"
	case licenseapi.Devsy:
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
