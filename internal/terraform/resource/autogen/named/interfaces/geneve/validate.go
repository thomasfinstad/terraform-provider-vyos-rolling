/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfacesgeneve code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesgeneve

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesGeneve{}
	_ resource.ResourceWithConfigure   = &interfacesGeneve{}
	_ resource.ResourceWithImportState = &interfacesGeneve{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesGeneve{}
// var _ resource.ResourceWithModifyPlan = &interfacesGeneve{}
// var _ resource.ResourceWithUpgradeState = &interfacesGeneve{}
// var _ resource.ResourceWithValidateConfig = &interfacesGeneve{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesGeneve{}
