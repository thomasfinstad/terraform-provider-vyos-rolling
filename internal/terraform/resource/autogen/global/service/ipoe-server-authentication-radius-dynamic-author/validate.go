// Package globalserviceipoeserverauthenticationradiusdynamicauthor code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceipoeserverauthenticationradiusdynamicauthor

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceIPoeServerAuthenticationRadiusDynamicAuthor{}
	_ resource.ResourceWithConfigure = &serviceIPoeServerAuthenticationRadiusDynamicAuthor{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerAuthenticationRadiusDynamicAuthor{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerAuthenticationRadiusDynamicAuthor{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerAuthenticationRadiusDynamicAuthor{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerAuthenticationRadiusDynamicAuthor{}
// var _ resource.ResourceWithImportState = &serviceIPoeServerAuthenticationRadiusDynamicAuthor{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerAuthenticationRadiusDynamicAuthor{}
