// Package namedvrfnameprotocolsospfinterfaceauthenticationmdfivekeyid code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfinterfaceauthenticationmdfivekeyid

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
