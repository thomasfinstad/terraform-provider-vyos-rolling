/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacestunnel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacestunnel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (tunnel) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesTunnel{}
	_ resource.ResourceWithConfigure   = &interfacesTunnel{}
	_ resource.ResourceWithImportState = &interfacesTunnel{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesTunnel{}
// var _ resource.ResourceWithModifyPlan = &interfacesTunnel{}
// var _ resource.ResourceWithUpgradeState = &interfacesTunnel{}
// var _ resource.ResourceWithValidateConfig = &interfacesTunnel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesTunnel{}
