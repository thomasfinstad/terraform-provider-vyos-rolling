/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsospfsegmentrouting code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfsegmentrouting

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfSegmentRouting{}
	_ resource.ResourceWithConfigure   = &protocolsOspfSegmentRouting{}
	_ resource.ResourceWithImportState = &protocolsOspfSegmentRouting{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfSegmentRouting{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfSegmentRouting{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfSegmentRouting{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfSegmentRouting{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfSegmentRouting{}
