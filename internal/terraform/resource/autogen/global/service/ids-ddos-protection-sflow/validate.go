/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalserviceidsddosprotectionsflow code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceidsddosprotectionsflow

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIDsDdosProtectionSflow{}
	_ resource.ResourceWithConfigure   = &serviceIDsDdosProtectionSflow{}
	_ resource.ResourceWithImportState = &serviceIDsDdosProtectionSflow{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIDsDdosProtectionSflow{}
// var _ resource.ResourceWithModifyPlan = &serviceIDsDdosProtectionSflow{}
// var _ resource.ResourceWithUpgradeState = &serviceIDsDdosProtectionSflow{}
// var _ resource.ResourceWithValidateConfig = &serviceIDsDdosProtectionSflow{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIDsDdosProtectionSflow{}
