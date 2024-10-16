// Package namedvrfnameprotocolsospfarea code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfarea

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsOspfArea{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsOspfArea{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsOspfArea{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfArea{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfArea{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfArea{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfArea{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfArea{}
