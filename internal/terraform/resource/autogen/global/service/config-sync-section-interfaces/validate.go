// Package globalserviceconfigsyncsectioninterfaces code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconfigsyncsectioninterfaces

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceConfigSyncSectionInterfaces{}
	_ resource.ResourceWithConfigure = &serviceConfigSyncSectionInterfaces{}
)

// var _ resource.ResourceWithConfigValidators = &serviceConfigSyncSectionInterfaces{}
// var _ resource.ResourceWithModifyPlan = &serviceConfigSyncSectionInterfaces{}
// var _ resource.ResourceWithUpgradeState = &serviceConfigSyncSectionInterfaces{}
// var _ resource.ResourceWithValidateConfig = &serviceConfigSyncSectionInterfaces{}
// var _ resource.ResourceWithImportState = &serviceConfigSyncSectionInterfaces{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceConfigSyncSectionInterfaces{}