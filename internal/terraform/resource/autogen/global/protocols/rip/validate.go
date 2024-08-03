// Package globalprotocolsrip code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsrip

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIP{}
	_ resource.ResourceWithConfigure = &protocolsRIP{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIP{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIP{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIP{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIP{}
// var _ resource.ResourceWithImportState = &protocolsRIP{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIP{}
