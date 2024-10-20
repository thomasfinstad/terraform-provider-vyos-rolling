/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfacespseudoethernetdhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetdhcpvsixoptionspdinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure   = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithImportState = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
