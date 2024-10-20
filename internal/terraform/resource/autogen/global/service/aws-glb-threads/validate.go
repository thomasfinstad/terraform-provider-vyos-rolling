/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalserviceawsglbthreads code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceawsglbthreads

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceAwsGlbThreads{}
	_ resource.ResourceWithConfigure   = &serviceAwsGlbThreads{}
	_ resource.ResourceWithImportState = &serviceAwsGlbThreads{}
)

// var _ resource.ResourceWithConfigValidators = &serviceAwsGlbThreads{}
// var _ resource.ResourceWithModifyPlan = &serviceAwsGlbThreads{}
// var _ resource.ResourceWithUpgradeState = &serviceAwsGlbThreads{}
// var _ resource.ResourceWithValidateConfig = &serviceAwsGlbThreads{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceAwsGlbThreads{}
