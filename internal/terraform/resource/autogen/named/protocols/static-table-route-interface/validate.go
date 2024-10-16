// Package namedprotocolsstatictablerouteinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictablerouteinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticTableRouteInterface{}
	_ resource.ResourceWithConfigure   = &protocolsStaticTableRouteInterface{}
	_ resource.ResourceWithImportState = &protocolsStaticTableRouteInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticTableRouteInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticTableRouteInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticTableRouteInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticTableRouteInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticTableRouteInterface{}
