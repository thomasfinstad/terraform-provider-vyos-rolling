/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicendpproxyinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicendpproxyinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceNdpProxyInterface{}
	_ resource.ResourceWithConfigure   = &serviceNdpProxyInterface{}
	_ resource.ResourceWithImportState = &serviceNdpProxyInterface{}
)

// var _ resource.ResourceWithConfigValidators = &serviceNdpProxyInterface{}
// var _ resource.ResourceWithModifyPlan = &serviceNdpProxyInterface{}
// var _ resource.ResourceWithUpgradeState = &serviceNdpProxyInterface{}
// var _ resource.ResourceWithValidateConfig = &serviceNdpProxyInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceNdpProxyInterface{}
