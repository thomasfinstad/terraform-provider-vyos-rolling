// Package globalprotocolspimsix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolspimsix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsPimsix{}
	_ resource.ResourceWithConfigure   = &protocolsPimsix{}
	_ resource.ResourceWithImportState = &protocolsPimsix{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsPimsix{}
// var _ resource.ResourceWithModifyPlan = &protocolsPimsix{}
// var _ resource.ResourceWithUpgradeState = &protocolsPimsix{}
// var _ resource.ResourceWithValidateConfig = &protocolsPimsix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsPimsix{}
