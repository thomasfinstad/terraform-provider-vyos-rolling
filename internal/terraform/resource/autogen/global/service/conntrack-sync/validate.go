/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalserviceconntracksync code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconntracksync

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceConntrackSync{}
	_ resource.ResourceWithConfigure   = &serviceConntrackSync{}
	_ resource.ResourceWithImportState = &serviceConntrackSync{}
)

// var _ resource.ResourceWithConfigValidators = &serviceConntrackSync{}
// var _ resource.ResourceWithModifyPlan = &serviceConntrackSync{}
// var _ resource.ResourceWithUpgradeState = &serviceConntrackSync{}
// var _ resource.ResourceWithValidateConfig = &serviceConntrackSync{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceConntrackSync{}
