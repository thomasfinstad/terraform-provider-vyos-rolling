// Package globalserviceconfigsyncsectionqos code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconfigsyncsectionqos

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceConfigSyncSectionQos{}
	_ resource.ResourceWithConfigure = &serviceConfigSyncSectionQos{}
)

// var _ resource.ResourceWithConfigValidators = &serviceConfigSyncSectionQos{}
// var _ resource.ResourceWithModifyPlan = &serviceConfigSyncSectionQos{}
// var _ resource.ResourceWithUpgradeState = &serviceConfigSyncSectionQos{}
// var _ resource.ResourceWithValidateConfig = &serviceConfigSyncSectionQos{}
// var _ resource.ResourceWithImportState = &serviceConfigSyncSectionQos{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceConfigSyncSectionQos{}
