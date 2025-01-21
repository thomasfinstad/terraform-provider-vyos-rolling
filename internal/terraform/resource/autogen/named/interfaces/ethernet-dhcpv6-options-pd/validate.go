/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacesethernetdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernetdhcpvsixoptionspd

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (pd) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesEthernetDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure   = &interfacesEthernetDhcpvsixOptionsPd{}
	_ resource.ResourceWithImportState = &interfacesEthernetDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesEthernetDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesEthernetDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesEthernetDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesEthernetDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesEthernetDhcpvsixOptionsPd{}
