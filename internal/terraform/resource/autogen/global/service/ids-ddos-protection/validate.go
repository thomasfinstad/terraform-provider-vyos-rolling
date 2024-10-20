/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalserviceidsddosprotection code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceidsddosprotection

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIDsDdosProtection{}
	_ resource.ResourceWithConfigure   = &serviceIDsDdosProtection{}
	_ resource.ResourceWithImportState = &serviceIDsDdosProtection{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIDsDdosProtection{}
// var _ resource.ResourceWithModifyPlan = &serviceIDsDdosProtection{}
// var _ resource.ResourceWithUpgradeState = &serviceIDsDdosProtection{}
// var _ resource.ResourceWithValidateConfig = &serviceIDsDdosProtection{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIDsDdosProtection{}
