// Package namedprotocolsospfinterfaceauthenticationmdfivekeyid code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfinterfaceauthenticationmdfivekeyid

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfInterfaceAuthenticationMdfiveKeyID{}
	_ resource.ResourceWithConfigure   = &protocolsOspfInterfaceAuthenticationMdfiveKeyID{}
	_ resource.ResourceWithImportState = &protocolsOspfInterfaceAuthenticationMdfiveKeyID{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfInterfaceAuthenticationMdfiveKeyID{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfInterfaceAuthenticationMdfiveKeyID{}
