/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicewebproxyurlfilteringsquidguardautoupdate code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicewebproxyurlfilteringsquidguardautoupdate

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (auto-update) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceWebproxyURLFilteringSquIDguardAutoUpdate{}
	_ resource.ResourceWithConfigure   = &serviceWebproxyURLFilteringSquIDguardAutoUpdate{}
	_ resource.ResourceWithImportState = &serviceWebproxyURLFilteringSquIDguardAutoUpdate{}
)

// var _ resource.ResourceWithConfigValidators = &serviceWebproxyURLFilteringSquIDguardAutoUpdate{}
// var _ resource.ResourceWithModifyPlan = &serviceWebproxyURLFilteringSquIDguardAutoUpdate{}
// var _ resource.ResourceWithUpgradeState = &serviceWebproxyURLFilteringSquIDguardAutoUpdate{}
// var _ resource.ResourceWithValidateConfig = &serviceWebproxyURLFilteringSquIDguardAutoUpdate{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceWebproxyURLFilteringSquIDguardAutoUpdate{}
