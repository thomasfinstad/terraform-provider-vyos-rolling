// Package namedvrfnameprotocolsstaticroutesix code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticroutesix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsStaticRoutesix{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsStaticRoutesix{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsStaticRoutesix{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsStaticRoutesix{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsStaticRoutesix{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsStaticRoutesix{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsStaticRoutesix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsStaticRoutesix{}
