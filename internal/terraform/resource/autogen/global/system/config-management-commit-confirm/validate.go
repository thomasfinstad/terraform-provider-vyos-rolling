// Package globalsystemconfigmanagementcommitconfirm code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemconfigmanagementcommitconfirm

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemConfigManagementCommitConfirm{}
	_ resource.ResourceWithConfigure = &systemConfigManagementCommitConfirm{}
)

// var _ resource.ResourceWithConfigValidators = &systemConfigManagementCommitConfirm{}
// var _ resource.ResourceWithModifyPlan = &systemConfigManagementCommitConfirm{}
// var _ resource.ResourceWithUpgradeState = &systemConfigManagementCommitConfirm{}
// var _ resource.ResourceWithValidateConfig = &systemConfigManagementCommitConfirm{}
// var _ resource.ResourceWithImportState = &systemConfigManagementCommitConfirm{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConfigManagementCommitConfirm{}