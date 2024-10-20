/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfacesethernetvifs code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernetvifs

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesEthernetVifS{}
	_ resource.ResourceWithConfigure   = &interfacesEthernetVifS{}
	_ resource.ResourceWithImportState = &interfacesEthernetVifS{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesEthernetVifS{}
// var _ resource.ResourceWithModifyPlan = &interfacesEthernetVifS{}
// var _ resource.ResourceWithUpgradeState = &interfacesEthernetVifS{}
// var _ resource.ResourceWithValidateConfig = &interfacesEthernetVifS{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesEthernetVifS{}
