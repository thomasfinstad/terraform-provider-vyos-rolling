// Package globalprotocolspimigmp code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolspimigmp

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsPimIgmp{}
	_ resource.ResourceWithConfigure = &protocolsPimIgmp{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsPimIgmp{}
// var _ resource.ResourceWithModifyPlan = &protocolsPimIgmp{}
// var _ resource.ResourceWithUpgradeState = &protocolsPimIgmp{}
// var _ resource.ResourceWithValidateConfig = &protocolsPimIgmp{}
// var _ resource.ResourceWithImportState = &protocolsPimIgmp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsPimIgmp{}