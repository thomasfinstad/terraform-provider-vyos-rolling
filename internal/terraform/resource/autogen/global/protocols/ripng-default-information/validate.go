// Package globalprotocolsripngdefaultinformation code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripngdefaultinformation

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPngDefaultInformation{}
	_ resource.ResourceWithConfigure   = &protocolsRIPngDefaultInformation{}
	_ resource.ResourceWithImportState = &protocolsRIPngDefaultInformation{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPngDefaultInformation{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPngDefaultInformation{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPngDefaultInformation{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPngDefaultInformation{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPngDefaultInformation{}
