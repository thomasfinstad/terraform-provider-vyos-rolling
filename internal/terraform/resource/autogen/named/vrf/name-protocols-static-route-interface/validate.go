// Package namedvrfnameprotocolsstaticrouteinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticrouteinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsStaticRouteInterface{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsStaticRouteInterface{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsStaticRouteInterface{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsStaticRouteInterface{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsStaticRouteInterface{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsStaticRouteInterface{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsStaticRouteInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsStaticRouteInterface{}
