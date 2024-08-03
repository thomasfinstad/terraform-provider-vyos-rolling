// Package globalsystemproxy code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemproxy

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemProxy{}
	_ resource.ResourceWithConfigure = &systemProxy{}
)

// var _ resource.ResourceWithConfigValidators = &systemProxy{}
// var _ resource.ResourceWithModifyPlan = &systemProxy{}
// var _ resource.ResourceWithUpgradeState = &systemProxy{}
// var _ resource.ResourceWithValidateConfig = &systemProxy{}
// var _ resource.ResourceWithImportState = &systemProxy{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemProxy{}
