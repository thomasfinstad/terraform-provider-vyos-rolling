/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalserviceipoeserverauthentication code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceipoeserverauthentication

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIPoeServerAuthentication{}
	_ resource.ResourceWithConfigure   = &serviceIPoeServerAuthentication{}
	_ resource.ResourceWithImportState = &serviceIPoeServerAuthentication{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerAuthentication{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerAuthentication{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerAuthentication{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerAuthentication{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerAuthentication{}
