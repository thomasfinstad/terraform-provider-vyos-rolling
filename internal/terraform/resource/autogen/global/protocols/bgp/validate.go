// Package globalprotocolsbgp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgp

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgp{}
	_ resource.ResourceWithConfigure   = &protocolsBgp{}
	_ resource.ResourceWithImportState = &protocolsBgp{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgp{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgp{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgp{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgp{}
