/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedsystemstatichostmappinghostname code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemstatichostmappinghostname

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (host-name) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemStaticHostMappingHostName{}
	_ resource.ResourceWithConfigure   = &systemStaticHostMappingHostName{}
	_ resource.ResourceWithImportState = &systemStaticHostMappingHostName{}
)

// var _ resource.ResourceWithConfigValidators = &systemStaticHostMappingHostName{}
// var _ resource.ResourceWithModifyPlan = &systemStaticHostMappingHostName{}
// var _ resource.ResourceWithUpgradeState = &systemStaticHostMappingHostName{}
// var _ resource.ResourceWithValidateConfig = &systemStaticHostMappingHostName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemStaticHostMappingHostName{}
