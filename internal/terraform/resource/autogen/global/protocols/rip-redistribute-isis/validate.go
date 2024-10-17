// Package globalprotocolsripredistributeisis code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripredistributeisis

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPRedistributeIsis{}
	_ resource.ResourceWithConfigure   = &protocolsRIPRedistributeIsis{}
	_ resource.ResourceWithImportState = &protocolsRIPRedistributeIsis{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPRedistributeIsis{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPRedistributeIsis{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPRedistributeIsis{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPRedistributeIsis{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPRedistributeIsis{}
