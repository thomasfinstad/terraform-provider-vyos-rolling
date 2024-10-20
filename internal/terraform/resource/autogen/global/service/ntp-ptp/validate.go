/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalservicentpptp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicentpptp

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceNtpPtp{}
	_ resource.ResourceWithConfigure   = &serviceNtpPtp{}
	_ resource.ResourceWithImportState = &serviceNtpPtp{}
)

// var _ resource.ResourceWithConfigValidators = &serviceNtpPtp{}
// var _ resource.ResourceWithModifyPlan = &serviceNtpPtp{}
// var _ resource.ResourceWithUpgradeState = &serviceNtpPtp{}
// var _ resource.ResourceWithValidateConfig = &serviceNtpPtp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceNtpPtp{}
