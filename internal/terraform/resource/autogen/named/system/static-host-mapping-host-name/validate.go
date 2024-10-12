// Package namedsystemstatichostmappinghostname code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemstatichostmappinghostname

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemStaticHostMappingHostName{}
	_ resource.ResourceWithConfigure = &systemStaticHostMappingHostName{}
)

// var _ resource.ResourceWithConfigValidators = &systemStaticHostMappingHostName{}
// var _ resource.ResourceWithModifyPlan = &systemStaticHostMappingHostName{}
// var _ resource.ResourceWithUpgradeState = &systemStaticHostMappingHostName{}
// var _ resource.ResourceWithValidateConfig = &systemStaticHostMappingHostName{}
// var _ resource.ResourceWithImportState = &systemStaticHostMappingHostName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemStaticHostMappingHostName{}
