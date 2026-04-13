package product

import (
	"github.com/skevetter/admin-apis/pkg/licenseapi"
)

// LoginCmd returns the login command for the product
func LoginCmd() string {
	loginCmd := "devsy login"

	switch Name() {
	case licenseapi.DevPodPro:
		return "devpod login"
	case licenseapi.DevsyPro:
		return "devsy platform login"
	case licenseapi.Devsy:
	}

	return loginCmd
}

// StartCmd returns the start command for the product
func StartCmd() string {
	loginCmd := "devsy start"

	switch Name() {
	case licenseapi.DevPodPro:
		loginCmd = "devpod pro start"
	case licenseapi.DevsyPro:
		loginCmd = "devsy platform start"
	case licenseapi.Devsy:
	}

	return loginCmd
}

// Url returns the url command for the product
func Url() string {
	loginCmd := "devsy-url"

	switch Name() {
	case licenseapi.DevPodPro:
		loginCmd = "devpod-pro-url"
	case licenseapi.DevsyPro:
		loginCmd = "devsy-pro-url"
	case licenseapi.Devsy:
	}

	return loginCmd
}

// ResetPassword returns the reset password command for the product
func ResetPassword() string {
	resetPassword := "devsy reset password"

	switch Name() {
	case licenseapi.DevPodPro:
		return "devpod pro reset password"
	case licenseapi.DevsyPro:
		return "devsy platform reset password"
	case licenseapi.Devsy:
	}

	return resetPassword
}
