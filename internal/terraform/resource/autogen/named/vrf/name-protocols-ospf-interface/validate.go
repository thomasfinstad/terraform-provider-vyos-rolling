// Package namedvrfnameprotocolsospfinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsOspfInterface{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsOspfInterface{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfInterface{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfInterface{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfInterface{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfInterface{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsOspfInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfInterface{}
