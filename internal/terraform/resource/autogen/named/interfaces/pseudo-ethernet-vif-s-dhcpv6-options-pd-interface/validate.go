/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacespseudoethernetvifsdhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetvifsdhcpvsixoptionspdinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (interface) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesPseudoEthernetVifSDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure   = &interfacesPseudoEthernetVifSDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithImportState = &interfacesPseudoEthernetVifSDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesPseudoEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesPseudoEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesPseudoEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesPseudoEthernetVifSDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesPseudoEthernetVifSDhcpvsixOptionsPdInterface{}
