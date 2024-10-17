// Package globalprotocolsisis code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisis

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsis{}
	_ resource.ResourceWithConfigure   = &protocolsIsis{}
	_ resource.ResourceWithImportState = &protocolsIsis{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsis{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsis{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsis{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsis{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsis{}
