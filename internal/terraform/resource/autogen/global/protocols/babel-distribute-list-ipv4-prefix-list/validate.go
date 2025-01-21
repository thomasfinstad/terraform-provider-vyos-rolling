/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbabeldistributelistipvfourprefixlist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbabeldistributelistipvfourprefixlist

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (prefix-list) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBabelDistributeListIPvfourPrefixList{}
	_ resource.ResourceWithConfigure   = &protocolsBabelDistributeListIPvfourPrefixList{}
	_ resource.ResourceWithImportState = &protocolsBabelDistributeListIPvfourPrefixList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBabelDistributeListIPvfourPrefixList{}
// var _ resource.ResourceWithModifyPlan = &protocolsBabelDistributeListIPvfourPrefixList{}
// var _ resource.ResourceWithUpgradeState = &protocolsBabelDistributeListIPvfourPrefixList{}
// var _ resource.ResourceWithValidateConfig = &protocolsBabelDistributeListIPvfourPrefixList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBabelDistributeListIPvfourPrefixList{}
