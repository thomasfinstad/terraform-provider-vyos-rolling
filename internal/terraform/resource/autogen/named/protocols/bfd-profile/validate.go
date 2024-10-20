/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedprotocolsbfdprofile code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbfdprofile

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBfdProfile{}
	_ resource.ResourceWithConfigure   = &protocolsBfdProfile{}
	_ resource.ResourceWithImportState = &protocolsBfdProfile{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBfdProfile{}
// var _ resource.ResourceWithModifyPlan = &protocolsBfdProfile{}
// var _ resource.ResourceWithUpgradeState = &protocolsBfdProfile{}
// var _ resource.ResourceWithValidateConfig = &protocolsBfdProfile{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBfdProfile{}
