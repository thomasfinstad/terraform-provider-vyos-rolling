// Package namedprotocolsospfaccesslist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfaccesslist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfAccessList{}
	_ resource.ResourceWithConfigure   = &protocolsOspfAccessList{}
	_ resource.ResourceWithImportState = &protocolsOspfAccessList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfAccessList{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfAccessList{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfAccessList{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfAccessList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfAccessList{}
