package main

import (
	"github.com/HashDataInc/terraform-provider-hashdata/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version = "0.0.1"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.New(version)})
}
