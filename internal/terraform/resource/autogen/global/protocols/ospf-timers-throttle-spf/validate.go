// Package globalprotocolsospftimersthrottlespf code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospftimersthrottlespf

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfTimersThroTTLeSpf{}
	_ resource.ResourceWithConfigure = &protocolsOspfTimersThroTTLeSpf{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfTimersThroTTLeSpf{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfTimersThroTTLeSpf{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfTimersThroTTLeSpf{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfTimersThroTTLeSpf{}
// var _ resource.ResourceWithImportState = &protocolsOspfTimersThroTTLeSpf{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfTimersThroTTLeSpf{}
