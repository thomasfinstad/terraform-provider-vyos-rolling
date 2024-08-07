// Package namedvrfnameprotocolsospfvthreeinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfvthreeinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsOspfvthreeInterface{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsOspfvthreeInterface{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfvthreeInterface{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfvthreeInterface{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfvthreeInterface{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfvthreeInterface{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsOspfvthreeInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfvthreeInterface{}
