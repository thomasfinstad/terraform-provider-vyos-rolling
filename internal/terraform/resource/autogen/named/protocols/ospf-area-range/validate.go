// Package namedprotocolsospfarearange code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfarearange

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfAreaRange{}
	_ resource.ResourceWithConfigure   = &protocolsOspfAreaRange{}
	_ resource.ResourceWithImportState = &protocolsOspfAreaRange{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfAreaRange{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfAreaRange{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfAreaRange{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfAreaRange{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfAreaRange{}
