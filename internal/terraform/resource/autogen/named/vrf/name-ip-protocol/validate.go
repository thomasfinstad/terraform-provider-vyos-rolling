/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameipprotocol code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameipprotocol

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (protocol) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameIPProtocol{}
	_ resource.ResourceWithConfigure   = &vrfNameIPProtocol{}
	_ resource.ResourceWithImportState = &vrfNameIPProtocol{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameIPProtocol{}
// var _ resource.ResourceWithModifyPlan = &vrfNameIPProtocol{}
// var _ resource.ResourceWithUpgradeState = &vrfNameIPProtocol{}
// var _ resource.ResourceWithValidateConfig = &vrfNameIPProtocol{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameIPProtocol{}
