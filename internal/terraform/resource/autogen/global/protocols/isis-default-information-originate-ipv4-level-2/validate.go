// Package globalprotocolsisisdefaultinformationoriginateipvfourleveltwo code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisdefaultinformationoriginateipvfourleveltwo

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisDefaultInformationOriginateIPvfourLevelTwo{}
	_ resource.ResourceWithConfigure = &protocolsIsisDefaultInformationOriginateIPvfourLevelTwo{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisDefaultInformationOriginateIPvfourLevelTwo{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisDefaultInformationOriginateIPvfourLevelTwo{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisDefaultInformationOriginateIPvfourLevelTwo{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisDefaultInformationOriginateIPvfourLevelTwo{}
// var _ resource.ResourceWithImportState = &protocolsIsisDefaultInformationOriginateIPvfourLevelTwo{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisDefaultInformationOriginateIPvfourLevelTwo{}
