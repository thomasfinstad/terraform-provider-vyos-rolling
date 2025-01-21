/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacesbondingdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbondingdhcpvsixoptionspd

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (pd) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesBondingDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure   = &interfacesBondingDhcpvsixOptionsPd{}
	_ resource.ResourceWithImportState = &interfacesBondingDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesBondingDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesBondingDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesBondingDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesBondingDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesBondingDhcpvsixOptionsPd{}
