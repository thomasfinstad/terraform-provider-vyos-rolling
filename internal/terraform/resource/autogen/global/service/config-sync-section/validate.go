// Package globalserviceconfigsyncsection code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconfigsyncsection

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceConfigSyncSection{}
	_ resource.ResourceWithConfigure   = &serviceConfigSyncSection{}
	_ resource.ResourceWithImportState = &serviceConfigSyncSection{}
)

// var _ resource.ResourceWithConfigValidators = &serviceConfigSyncSection{}
// var _ resource.ResourceWithModifyPlan = &serviceConfigSyncSection{}
// var _ resource.ResourceWithUpgradeState = &serviceConfigSyncSection{}
// var _ resource.ResourceWithValidateConfig = &serviceConfigSyncSection{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceConfigSyncSection{}
