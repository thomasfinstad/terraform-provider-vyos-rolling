// Package namedqospolicyshaper code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaper

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyShaper{}
	_ resource.ResourceWithConfigure = &qosPolicyShaper{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyShaper{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyShaper{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyShaper{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyShaper{}
// var _ resource.ResourceWithImportState = &qosPolicyShaper{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyShaper{}
