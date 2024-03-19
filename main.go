package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider"
)

// Run "go generate" to generate the docs for the registry/website

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate tfplugindocs generate --provider-name vyos

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
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

// TODO Clean up logging
//  Find out if there is a better way to allow logging during tests
//  and general terraform logging to cooperate. Also clean up different
//  logging levels being used and create a clear guideline for what
//  is included in the different log levels.

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
//  - [ ] Investigate how to add "chglog" friendly part to commit message
//  - [ ] Add git commit messages
//  - [ ] Autogenerate resource changes from provider-schema files
