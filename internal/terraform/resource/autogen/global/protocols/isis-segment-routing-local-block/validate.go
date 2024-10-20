/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsisissegmentroutinglocalblock code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisissegmentroutinglocalblock

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisSegmentRoutingLocalBlock{}
	_ resource.ResourceWithConfigure   = &protocolsIsisSegmentRoutingLocalBlock{}
	_ resource.ResourceWithImportState = &protocolsIsisSegmentRoutingLocalBlock{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisSegmentRoutingLocalBlock{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisSegmentRoutingLocalBlock{}
