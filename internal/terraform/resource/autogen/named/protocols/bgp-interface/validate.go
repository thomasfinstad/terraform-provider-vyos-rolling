/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsbgpinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (interface) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpInterface{}
	_ resource.ResourceWithConfigure   = &protocolsBgpInterface{}
	_ resource.ResourceWithImportState = &protocolsBgpInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpInterface{}
