/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicehttpscertificates code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicehttpscertificates

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (certificates) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceHTTPSCertificates{}
	_ resource.ResourceWithConfigure   = &serviceHTTPSCertificates{}
	_ resource.ResourceWithImportState = &serviceHTTPSCertificates{}
)

// var _ resource.ResourceWithConfigValidators = &serviceHTTPSCertificates{}
// var _ resource.ResourceWithModifyPlan = &serviceHTTPSCertificates{}
// var _ resource.ResourceWithUpgradeState = &serviceHTTPSCertificates{}
// var _ resource.ResourceWithValidateConfig = &serviceHTTPSCertificates{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceHTTPSCertificates{}
