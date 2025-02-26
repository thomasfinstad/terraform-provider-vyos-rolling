/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfaceswirelessvifs code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvifs

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (vif-s) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesWirelessVifS{}
	_ resource.ResourceWithConfigure   = &interfacesWirelessVifS{}
	_ resource.ResourceWithImportState = &interfacesWirelessVifS{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWirelessVifS{}
// var _ resource.ResourceWithModifyPlan = &interfacesWirelessVifS{}
// var _ resource.ResourceWithUpgradeState = &interfacesWirelessVifS{}
// var _ resource.ResourceWithValidateConfig = &interfacesWirelessVifS{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWirelessVifS{}
