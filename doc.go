// Package credentialhelper helps create a Terraform credential helper CLI.
//
// https://www.terraform.io/docs/internals/credentials-helpers.html
//
// It accepts a "Helper" interface that implements the get, store, and forget command
// that Terraform will call. It prints output (or errors) to the correct output stream
// and sets an exit code based on the result of the Helper. It also supports defining flags,
// to be provided when the user defines their credential helper:
//
//   credentials_helper "credstore" {
//     args = ["--host=credstore.example.com"]
//   }
package credentialhelper
