// Package globalsystemsflow code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemsflow

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemSflow{}
	_ resource.ResourceWithConfigure = &systemSflow{}
)

// var _ resource.ResourceWithConfigValidators = &systemSflow{}
// var _ resource.ResourceWithModifyPlan = &systemSflow{}
// var _ resource.ResourceWithUpgradeState = &systemSflow{}
// var _ resource.ResourceWithValidateConfig = &systemSflow{}
// var _ resource.ResourceWithImportState = &systemSflow{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemSflow{}