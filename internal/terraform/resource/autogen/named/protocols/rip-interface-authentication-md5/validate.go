// Package namedprotocolsripinterfaceauthenticationmdfive code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsripinterfaceauthenticationmdfive

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPInterfaceAuthenticationMdfive{}
	_ resource.ResourceWithConfigure = &protocolsRIPInterfaceAuthenticationMdfive{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPInterfaceAuthenticationMdfive{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPInterfaceAuthenticationMdfive{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPInterfaceAuthenticationMdfive{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPInterfaceAuthenticationMdfive{}
// var _ resource.ResourceWithImportState = &protocolsRIPInterfaceAuthenticationMdfive{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPInterfaceAuthenticationMdfive{}
