/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsospfsegmentroutinglocalblock code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfsegmentroutinglocalblock

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfSegmentRoutingLocalBlock{}
	_ resource.ResourceWithConfigure   = &protocolsOspfSegmentRoutingLocalBlock{}
	_ resource.ResourceWithImportState = &protocolsOspfSegmentRoutingLocalBlock{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfSegmentRoutingLocalBlock{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfSegmentRoutingLocalBlock{}
