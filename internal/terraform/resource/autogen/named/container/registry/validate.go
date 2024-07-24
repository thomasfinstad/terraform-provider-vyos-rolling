// Package namedcontainerregistry code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedcontainerregistry

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &containerRegistry{}
	_ resource.ResourceWithConfigure = &containerRegistry{}
)

// var _ resource.ResourceWithConfigValidators = &containerRegistry{}
// var _ resource.ResourceWithModifyPlan = &containerRegistry{}
// var _ resource.ResourceWithUpgradeState = &containerRegistry{}
// var _ resource.ResourceWithValidateConfig = &containerRegistry{}
// var _ resource.ResourceWithImportState = &containerRegistry{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &containerRegistry{}
