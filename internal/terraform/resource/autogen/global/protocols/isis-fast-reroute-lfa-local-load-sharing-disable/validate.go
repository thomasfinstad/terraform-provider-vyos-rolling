// Package globalprotocolsisisfastreroutelfalocalloadsharingdisable code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisfastreroutelfalocalloadsharingdisable

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisFastRerouteLfaLocalLoadSharingDisable{}
	_ resource.ResourceWithConfigure = &protocolsIsisFastRerouteLfaLocalLoadSharingDisable{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisFastRerouteLfaLocalLoadSharingDisable{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisFastRerouteLfaLocalLoadSharingDisable{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisFastRerouteLfaLocalLoadSharingDisable{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisFastRerouteLfaLocalLoadSharingDisable{}
// var _ resource.ResourceWithImportState = &protocolsIsisFastRerouteLfaLocalLoadSharingDisable{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisFastRerouteLfaLocalLoadSharingDisable{}
