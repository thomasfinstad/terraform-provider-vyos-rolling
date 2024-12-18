/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfaceswireguardpeer code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswireguardpeer

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesWireguardPeer{}
	_ resource.ResourceWithConfigure   = &interfacesWireguardPeer{}
	_ resource.ResourceWithImportState = &interfacesWireguardPeer{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWireguardPeer{}
// var _ resource.ResourceWithModifyPlan = &interfacesWireguardPeer{}
// var _ resource.ResourceWithUpgradeState = &interfacesWireguardPeer{}
// var _ resource.ResourceWithValidateConfig = &interfacesWireguardPeer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWireguardPeer{}
