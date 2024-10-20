/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfaceswirelessvifsvifcdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvifsvifcdhcpvsixoptionspd

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure   = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
	_ resource.ResourceWithImportState = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
