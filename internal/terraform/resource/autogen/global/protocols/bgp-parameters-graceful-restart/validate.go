// Package globalprotocolsbgpparametersgracefulrestart code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpparametersgracefulrestart

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpParametersGracefulRestart{}
	_ resource.ResourceWithConfigure = &protocolsBgpParametersGracefulRestart{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpParametersGracefulRestart{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpParametersGracefulRestart{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpParametersGracefulRestart{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpParametersGracefulRestart{}
// var _ resource.ResourceWithImportState = &protocolsBgpParametersGracefulRestart{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpParametersGracefulRestart{}
