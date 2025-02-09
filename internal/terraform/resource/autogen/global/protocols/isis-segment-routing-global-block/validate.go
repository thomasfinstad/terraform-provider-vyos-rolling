/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsisissegmentroutingglobalblock code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisissegmentroutingglobalblock

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (global-block) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisSegmentRoutingGlobalBlock{}
	_ resource.ResourceWithConfigure   = &protocolsIsisSegmentRoutingGlobalBlock{}
	_ resource.ResourceWithImportState = &protocolsIsisSegmentRoutingGlobalBlock{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisSegmentRoutingGlobalBlock{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisSegmentRoutingGlobalBlock{}
