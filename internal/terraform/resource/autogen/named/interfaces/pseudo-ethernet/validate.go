/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfacespseudoethernet code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernet

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesPseudoEthernet{}
	_ resource.ResourceWithConfigure   = &interfacesPseudoEthernet{}
	_ resource.ResourceWithImportState = &interfacesPseudoEthernet{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesPseudoEthernet{}
// var _ resource.ResourceWithModifyPlan = &interfacesPseudoEthernet{}
// var _ resource.ResourceWithUpgradeState = &interfacesPseudoEthernet{}
// var _ resource.ResourceWithValidateConfig = &interfacesPseudoEthernet{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesPseudoEthernet{}
