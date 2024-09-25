// Package globalprotocolsospfvthreedistance code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfvthreedistance

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfvthreeDistance{}
	_ resource.ResourceWithConfigure = &protocolsOspfvthreeDistance{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfvthreeDistance{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfvthreeDistance{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfvthreeDistance{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfvthreeDistance{}
// var _ resource.ResourceWithImportState = &protocolsOspfvthreeDistance{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfvthreeDistance{}
