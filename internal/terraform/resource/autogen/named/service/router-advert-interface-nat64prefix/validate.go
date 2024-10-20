/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicerouteradvertinterfacenatsixfourprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicerouteradvertinterfacenatsixfourprefix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceRouterAdvertInterfaceNatsixfourprefix{}
	_ resource.ResourceWithConfigure   = &serviceRouterAdvertInterfaceNatsixfourprefix{}
	_ resource.ResourceWithImportState = &serviceRouterAdvertInterfaceNatsixfourprefix{}
)

// var _ resource.ResourceWithConfigValidators = &serviceRouterAdvertInterfaceNatsixfourprefix{}
// var _ resource.ResourceWithModifyPlan = &serviceRouterAdvertInterfaceNatsixfourprefix{}
// var _ resource.ResourceWithUpgradeState = &serviceRouterAdvertInterfaceNatsixfourprefix{}
// var _ resource.ResourceWithValidateConfig = &serviceRouterAdvertInterfaceNatsixfourprefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceRouterAdvertInterfaceNatsixfourprefix{}
