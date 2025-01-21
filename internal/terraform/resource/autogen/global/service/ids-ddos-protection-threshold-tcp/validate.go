/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalserviceidsddosprotectionthresholdtcp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceidsddosprotectionthresholdtcp

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (tcp) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIDsDdosProtectionThresholdTCP{}
	_ resource.ResourceWithConfigure   = &serviceIDsDdosProtectionThresholdTCP{}
	_ resource.ResourceWithImportState = &serviceIDsDdosProtectionThresholdTCP{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIDsDdosProtectionThresholdTCP{}
// var _ resource.ResourceWithModifyPlan = &serviceIDsDdosProtectionThresholdTCP{}
// var _ resource.ResourceWithUpgradeState = &serviceIDsDdosProtectionThresholdTCP{}
// var _ resource.ResourceWithValidateConfig = &serviceIDsDdosProtectionThresholdTCP{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIDsDdosProtectionThresholdTCP{}
