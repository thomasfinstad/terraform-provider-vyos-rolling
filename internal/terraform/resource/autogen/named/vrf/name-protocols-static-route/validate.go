// Package namedvrfnameprotocolsstaticroute code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticroute

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsStaticRoute{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsStaticRoute{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsStaticRoute{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsStaticRoute{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsStaticRoute{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsStaticRoute{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsStaticRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsStaticRoute{}
