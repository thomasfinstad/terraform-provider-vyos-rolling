// Package globalprotocolsisisdefaultinformationoriginateipvsixlevelone code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisdefaultinformationoriginateipvsixlevelone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisDefaultInformationOriginateIPvsixLevelOne{}
	_ resource.ResourceWithConfigure = &protocolsIsisDefaultInformationOriginateIPvsixLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisDefaultInformationOriginateIPvsixLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisDefaultInformationOriginateIPvsixLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisDefaultInformationOriginateIPvsixLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisDefaultInformationOriginateIPvsixLevelOne{}
// var _ resource.ResourceWithImportState = &protocolsIsisDefaultInformationOriginateIPvsixLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisDefaultInformationOriginateIPvsixLevelOne{}
