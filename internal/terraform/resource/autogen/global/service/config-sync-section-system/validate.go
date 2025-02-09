/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalserviceconfigsyncsectionsystem code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconfigsyncsectionsystem

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (system) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceConfigSyncSectionSystem{}
	_ resource.ResourceWithConfigure   = &serviceConfigSyncSectionSystem{}
	_ resource.ResourceWithImportState = &serviceConfigSyncSectionSystem{}
)

// var _ resource.ResourceWithConfigValidators = &serviceConfigSyncSectionSystem{}
// var _ resource.ResourceWithModifyPlan = &serviceConfigSyncSectionSystem{}
// var _ resource.ResourceWithUpgradeState = &serviceConfigSyncSectionSystem{}
// var _ resource.ResourceWithValidateConfig = &serviceConfigSyncSectionSystem{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceConfigSyncSectionSystem{}
