/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfnameprotocolsospfsegmentroutingprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfsegmentroutingprefix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfSegmentRoutingPrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
