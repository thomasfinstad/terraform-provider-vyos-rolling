/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalserviceidsddosprotectionthresholdicmp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceidsddosprotectionthresholdicmp

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIDsDdosProtectionThresholdIcmp{}
	_ resource.ResourceWithConfigure   = &serviceIDsDdosProtectionThresholdIcmp{}
	_ resource.ResourceWithImportState = &serviceIDsDdosProtectionThresholdIcmp{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIDsDdosProtectionThresholdIcmp{}
// var _ resource.ResourceWithModifyPlan = &serviceIDsDdosProtectionThresholdIcmp{}
// var _ resource.ResourceWithUpgradeState = &serviceIDsDdosProtectionThresholdIcmp{}
// var _ resource.ResourceWithValidateConfig = &serviceIDsDdosProtectionThresholdIcmp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIDsDdosProtectionThresholdIcmp{}
