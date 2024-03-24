package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider"
)

var (
	// TODO set correct version
	//  milestone:1

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
		//  milestone:1
		Address: "github.com/thomasfinstad/vyos",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}

// TODO Create a CONTRIBUTION.md doc
//  milestone:2
//  include meta info such as
//  - [ ] how schema is used
//  - [ ] how the code generation works
//  - [ ] why so many extra files are committed to the repo
//  - [ ] diagram of the makefile workflow

// TODO Autogenerate CHANGELOG.md
//  milestone:2
//  Ref: https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#versioning-specification
//  Ref: https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#changelog-specification
//  - [ ] Investigate how to add "chglog" friendly part to commit message
//  - [ ] Add git commit messages
//  - [ ] Autogenerate resource changes from provider-schema files

// TODO Investigate official code gen method
//  Ref: https://developer.hashicorp.com/terraform/plugin/code-generation/design
//  milestone:6
