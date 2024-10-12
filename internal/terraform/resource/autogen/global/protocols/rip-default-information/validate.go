// Package globalprotocolsripdefaultinformation code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripdefaultinformation

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPDefaultInformation{}
	_ resource.ResourceWithConfigure = &protocolsRIPDefaultInformation{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPDefaultInformation{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPDefaultInformation{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPDefaultInformation{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPDefaultInformation{}
// var _ resource.ResourceWithImportState = &protocolsRIPDefaultInformation{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPDefaultInformation{}
