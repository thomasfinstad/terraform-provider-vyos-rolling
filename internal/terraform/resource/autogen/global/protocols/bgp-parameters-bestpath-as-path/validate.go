// Package globalprotocolsbgpparametersbestpathaspath code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpparametersbestpathaspath

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpParametersBestpathAsPath{}
	_ resource.ResourceWithConfigure   = &protocolsBgpParametersBestpathAsPath{}
	_ resource.ResourceWithImportState = &protocolsBgpParametersBestpathAsPath{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpParametersBestpathAsPath{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpParametersBestpathAsPath{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpParametersBestpathAsPath{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpParametersBestpathAsPath{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpParametersBestpathAsPath{}
