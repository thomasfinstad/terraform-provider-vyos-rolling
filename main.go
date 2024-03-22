package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	// TODO set correct version
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// TODO: Update this string with the published name of provider.
		Address: "github.com/thomasfinstad/vyos",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}

// TODO Mask any known sensitive information
//  Such as API keys. Also investigate if there is any information
//  in the vyos schema about (leaf)nodes that contain sensitive
//  information.

// TODO Create a CONTRIBUTION.md doc
//  include meta info such as
//  - [ ] how schema is used
//  - [ ] how the code generation works
//  - [ ] why so many extra files are commited to the repo
//  - [ ] diagram of the makefile workflow

// TODO Autogenerate CHANGELOG.md
//  Ref: https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#changelog-specification
//  - [ ] Investigate how to add "chglog" friendly part to commit message
//  - [ ] Add git commit messages
//  - [ ] Autogenerate resource changes from provider-schema files
