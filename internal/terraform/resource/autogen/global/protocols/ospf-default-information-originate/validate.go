// Package globalprotocolsospfdefaultinformationoriginate code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfdefaultinformationoriginate

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfDefaultInformationOriginate{}
	_ resource.ResourceWithConfigure = &protocolsOspfDefaultInformationOriginate{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfDefaultInformationOriginate{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfDefaultInformationOriginate{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfDefaultInformationOriginate{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfDefaultInformationOriginate{}
// var _ resource.ResourceWithImportState = &protocolsOspfDefaultInformationOriginate{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfDefaultInformationOriginate{}
