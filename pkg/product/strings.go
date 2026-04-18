package product

import (
	"github.com/devsy-org/admin-apis/pkg/licenseapi"
)

// LoginCmd returns the login command for the product
func LoginCmd() string {
	switch Name() {
	case licenseapi.DevsyPro:
		return "devsy platform login"
	case licenseapi.DevsyOrg:
		return "devsy login"
	}

	return "devsy login"
}

// StartCmd returns the start command for the product
func StartCmd() string {
	switch Name() {
	case licenseapi.DevsyPro:
		return "devsy platform start"
	case licenseapi.DevsyOrg:
		return "devsy start"
	}

	return "devsy start"
}

// Url returns the url command for the product
func Url() string {
	switch Name() {
	case licenseapi.DevsyPro:
		return "devsy-pro-url"
	case licenseapi.DevsyOrg:
		return "devsy-url"
	}

	return "devsy-url"
}

// ResetPassword returns the reset password command for the product
func ResetPassword() string {
	switch Name() {
	case licenseapi.DevsyPro:
		return "devsy platform reset password"
	case licenseapi.DevsyOrg:
		return "devsy reset password"
	}

	return "devsy reset password"
}
