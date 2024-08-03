// Package globalprotocolsmplsparameters code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsparameters

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsMplsParameters{}
	_ resource.ResourceWithConfigure = &protocolsMplsParameters{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsMplsParameters{}
// var _ resource.ResourceWithModifyPlan = &protocolsMplsParameters{}
// var _ resource.ResourceWithUpgradeState = &protocolsMplsParameters{}
// var _ resource.ResourceWithValidateConfig = &protocolsMplsParameters{}
// var _ resource.ResourceWithImportState = &protocolsMplsParameters{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsMplsParameters{}
