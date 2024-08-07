// Package namedvrfnameprotocolsospfaccesslist code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfaccesslist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsOspfAccessList{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsOspfAccessList{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfAccessList{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfAccessList{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfAccessList{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfAccessList{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsOspfAccessList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfAccessList{}
