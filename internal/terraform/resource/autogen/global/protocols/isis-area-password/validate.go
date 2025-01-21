/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsisisareapassword code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisareapassword

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (area-password) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisAreaPassword{}
	_ resource.ResourceWithConfigure   = &protocolsIsisAreaPassword{}
	_ resource.ResourceWithImportState = &protocolsIsisAreaPassword{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisAreaPassword{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisAreaPassword{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisAreaPassword{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisAreaPassword{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisAreaPassword{}
