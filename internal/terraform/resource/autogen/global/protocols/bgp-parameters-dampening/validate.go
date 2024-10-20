/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsbgpparametersdampening code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpparametersdampening

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpParametersDampening{}
	_ resource.ResourceWithConfigure   = &protocolsBgpParametersDampening{}
	_ resource.ResourceWithImportState = &protocolsBgpParametersDampening{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpParametersDampening{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpParametersDampening{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpParametersDampening{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpParametersDampening{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpParametersDampening{}
