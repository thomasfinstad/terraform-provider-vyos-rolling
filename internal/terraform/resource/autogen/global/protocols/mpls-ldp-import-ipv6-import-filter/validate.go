// Package globalprotocolsmplsldpimportipvsiximportfilter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsldpimportipvsiximportfilter

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsMplsLdpImportIPvsixImportFilter{}
	_ resource.ResourceWithConfigure = &protocolsMplsLdpImportIPvsixImportFilter{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsMplsLdpImportIPvsixImportFilter{}
// var _ resource.ResourceWithModifyPlan = &protocolsMplsLdpImportIPvsixImportFilter{}
// var _ resource.ResourceWithUpgradeState = &protocolsMplsLdpImportIPvsixImportFilter{}
// var _ resource.ResourceWithValidateConfig = &protocolsMplsLdpImportIPvsixImportFilter{}
// var _ resource.ResourceWithImportState = &protocolsMplsLdpImportIPvsixImportFilter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsMplsLdpImportIPvsixImportFilter{}
