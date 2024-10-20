/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedserviceconntracksyncinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceconntracksyncinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceConntrackSyncInterface{}
	_ resource.ResourceWithConfigure   = &serviceConntrackSyncInterface{}
	_ resource.ResourceWithImportState = &serviceConntrackSyncInterface{}
)

// var _ resource.ResourceWithConfigValidators = &serviceConntrackSyncInterface{}
// var _ resource.ResourceWithModifyPlan = &serviceConntrackSyncInterface{}
// var _ resource.ResourceWithUpgradeState = &serviceConntrackSyncInterface{}
// var _ resource.ResourceWithValidateConfig = &serviceConntrackSyncInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceConntrackSyncInterface{}
