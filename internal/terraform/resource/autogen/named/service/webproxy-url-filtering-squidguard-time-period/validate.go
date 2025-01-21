/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicewebproxyurlfilteringsquidguardtimeperiod code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicewebproxyurlfilteringsquidguardtimeperiod

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (time-period) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceWebproxyURLFilteringSquIDguardTimePeriod{}
	_ resource.ResourceWithConfigure   = &serviceWebproxyURLFilteringSquIDguardTimePeriod{}
	_ resource.ResourceWithImportState = &serviceWebproxyURLFilteringSquIDguardTimePeriod{}
)

// var _ resource.ResourceWithConfigValidators = &serviceWebproxyURLFilteringSquIDguardTimePeriod{}
// var _ resource.ResourceWithModifyPlan = &serviceWebproxyURLFilteringSquIDguardTimePeriod{}
// var _ resource.ResourceWithUpgradeState = &serviceWebproxyURLFilteringSquIDguardTimePeriod{}
// var _ resource.ResourceWithValidateConfig = &serviceWebproxyURLFilteringSquIDguardTimePeriod{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceWebproxyURLFilteringSquIDguardTimePeriod{}
