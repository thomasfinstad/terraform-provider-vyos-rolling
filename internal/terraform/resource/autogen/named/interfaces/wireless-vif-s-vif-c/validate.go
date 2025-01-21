/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfaceswirelessvifsvifc code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvifsvifc

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (vif-c) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesWirelessVifSVifC{}
	_ resource.ResourceWithConfigure   = &interfacesWirelessVifSVifC{}
	_ resource.ResourceWithImportState = &interfacesWirelessVifSVifC{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWirelessVifSVifC{}
// var _ resource.ResourceWithModifyPlan = &interfacesWirelessVifSVifC{}
// var _ resource.ResourceWithUpgradeState = &interfacesWirelessVifSVifC{}
// var _ resource.ResourceWithValidateConfig = &interfacesWirelessVifSVifC{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWirelessVifSVifC{}
