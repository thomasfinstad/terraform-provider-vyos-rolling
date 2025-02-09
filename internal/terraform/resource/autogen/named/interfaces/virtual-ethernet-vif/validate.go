/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacesvirtualethernetvif code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvirtualethernetvif

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (vif) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesVirtualEthernetVif{}
	_ resource.ResourceWithConfigure   = &interfacesVirtualEthernetVif{}
	_ resource.ResourceWithImportState = &interfacesVirtualEthernetVif{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesVirtualEthernetVif{}
// var _ resource.ResourceWithModifyPlan = &interfacesVirtualEthernetVif{}
// var _ resource.ResourceWithUpgradeState = &interfacesVirtualEthernetVif{}
// var _ resource.ResourceWithValidateConfig = &interfacesVirtualEthernetVif{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesVirtualEthernetVif{}
