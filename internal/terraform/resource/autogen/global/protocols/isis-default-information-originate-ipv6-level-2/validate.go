// Package globalprotocolsisisdefaultinformationoriginateipvsixleveltwo code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisdefaultinformationoriginateipvsixleveltwo

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}
	_ resource.ResourceWithConfigure   = &protocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}
	_ resource.ResourceWithImportState = &protocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}
