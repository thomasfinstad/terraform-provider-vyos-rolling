// Package globalprotocolsisisdefaultinformationoriginateipvfourlevelone code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisdefaultinformationoriginateipvfourlevelone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisDefaultInformationOriginateIPvfourLevelOne{}
	_ resource.ResourceWithConfigure   = &protocolsIsisDefaultInformationOriginateIPvfourLevelOne{}
	_ resource.ResourceWithImportState = &protocolsIsisDefaultInformationOriginateIPvfourLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisDefaultInformationOriginateIPvfourLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisDefaultInformationOriginateIPvfourLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisDefaultInformationOriginateIPvfourLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisDefaultInformationOriginateIPvfourLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisDefaultInformationOriginateIPvfourLevelOne{}
