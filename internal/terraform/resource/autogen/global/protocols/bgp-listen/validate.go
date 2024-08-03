// Package globalprotocolsbgplisten code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgplisten

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpListen{}
	_ resource.ResourceWithConfigure = &protocolsBgpListen{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpListen{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpListen{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpListen{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpListen{}
// var _ resource.ResourceWithImportState = &protocolsBgpListen{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpListen{}
