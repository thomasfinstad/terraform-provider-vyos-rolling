/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedserviceipoeserverauthenticationinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverauthenticationinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIPoeServerAuthenticationInterface{}
	_ resource.ResourceWithConfigure   = &serviceIPoeServerAuthenticationInterface{}
	_ resource.ResourceWithImportState = &serviceIPoeServerAuthenticationInterface{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerAuthenticationInterface{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerAuthenticationInterface{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerAuthenticationInterface{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerAuthenticationInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerAuthenticationInterface{}
