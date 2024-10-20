/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedloadbalancinghaproxyserviceloggingfacility code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancinghaproxyserviceloggingfacility

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &loadBalancingHaproxyServiceLoggingFacility{}
	_ resource.ResourceWithConfigure   = &loadBalancingHaproxyServiceLoggingFacility{}
	_ resource.ResourceWithImportState = &loadBalancingHaproxyServiceLoggingFacility{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingHaproxyServiceLoggingFacility{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingHaproxyServiceLoggingFacility{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingHaproxyServiceLoggingFacility{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingHaproxyServiceLoggingFacility{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingHaproxyServiceLoggingFacility{}
