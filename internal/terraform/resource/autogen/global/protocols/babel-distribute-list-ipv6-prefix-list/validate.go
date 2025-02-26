/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbabeldistributelistipvsixprefixlist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbabeldistributelistipvsixprefixlist

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (prefix-list) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBabelDistributeListIPvsixPrefixList{}
	_ resource.ResourceWithConfigure   = &protocolsBabelDistributeListIPvsixPrefixList{}
	_ resource.ResourceWithImportState = &protocolsBabelDistributeListIPvsixPrefixList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBabelDistributeListIPvsixPrefixList{}
// var _ resource.ResourceWithModifyPlan = &protocolsBabelDistributeListIPvsixPrefixList{}
// var _ resource.ResourceWithUpgradeState = &protocolsBabelDistributeListIPvsixPrefixList{}
// var _ resource.ResourceWithValidateConfig = &protocolsBabelDistributeListIPvsixPrefixList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBabelDistributeListIPvsixPrefixList{}
