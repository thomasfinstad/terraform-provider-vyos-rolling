// Package namedprotocolsisisfastreroutelfaremoteprefixlist code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsisisfastreroutelfaremoteprefixlist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisFastRerouteLfaRemotePrefixList{}
	_ resource.ResourceWithConfigure = &protocolsIsisFastRerouteLfaRemotePrefixList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisFastRerouteLfaRemotePrefixList{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisFastRerouteLfaRemotePrefixList{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisFastRerouteLfaRemotePrefixList{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisFastRerouteLfaRemotePrefixList{}
// var _ resource.ResourceWithImportState = &protocolsIsisFastRerouteLfaRemotePrefixList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisFastRerouteLfaRemotePrefixList{}
