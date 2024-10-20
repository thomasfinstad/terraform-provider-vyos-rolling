/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsrpki code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsrpki

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRpki{}
	_ resource.ResourceWithConfigure   = &protocolsRpki{}
	_ resource.ResourceWithImportState = &protocolsRpki{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRpki{}
// var _ resource.ResourceWithModifyPlan = &protocolsRpki{}
// var _ resource.ResourceWithUpgradeState = &protocolsRpki{}
// var _ resource.ResourceWithValidateConfig = &protocolsRpki{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRpki{}
