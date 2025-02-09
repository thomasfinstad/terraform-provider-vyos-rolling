/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedserviceipoeserverauthenticationradiusserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverauthenticationradiusserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (server) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIPoeServerAuthenticationRadiusServer{}
	_ resource.ResourceWithConfigure   = &serviceIPoeServerAuthenticationRadiusServer{}
	_ resource.ResourceWithImportState = &serviceIPoeServerAuthenticationRadiusServer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerAuthenticationRadiusServer{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerAuthenticationRadiusServer{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerAuthenticationRadiusServer{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerAuthenticationRadiusServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerAuthenticationRadiusServer{}
