// Package globalsystemlogintacacs code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemlogintacacs

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemLoginTacacs{}
	_ resource.ResourceWithConfigure = &systemLoginTacacs{}
)

// var _ resource.ResourceWithConfigValidators = &systemLoginTacacs{}
// var _ resource.ResourceWithModifyPlan = &systemLoginTacacs{}
// var _ resource.ResourceWithUpgradeState = &systemLoginTacacs{}
// var _ resource.ResourceWithValidateConfig = &systemLoginTacacs{}
// var _ resource.ResourceWithImportState = &systemLoginTacacs{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemLoginTacacs{}