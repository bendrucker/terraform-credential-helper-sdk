package credentialhelper

import "flag"

// Helper represents a Terraform credential helper
// https://www.terraform.io/docs/internals/credentials-helpers.html
type Helper interface {
	// Retrieve the credentials for the given hostname
	Get(hostname string, flags *flag.FlagSet) (credentials []byte, err error)
	// Store new credentials for the given hostname
	Store(hostname string, credentials []byte, flags *flag.FlagSet) error
	// Delete any stored credentials for the given hostname
	Forget(hostname string, flags *flag.FlagSet) error
}

//go:generate mockgen -destination=helper_test.go -package credentialhelper github.com/bendrucker/terraform-credential-helper-sdk Helper
