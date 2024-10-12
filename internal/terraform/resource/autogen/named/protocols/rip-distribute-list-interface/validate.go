// Package namedprotocolsripdistributelistinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsripdistributelistinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPDistributeListInterface{}
	_ resource.ResourceWithConfigure = &protocolsRIPDistributeListInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPDistributeListInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPDistributeListInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPDistributeListInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPDistributeListInterface{}
// var _ resource.ResourceWithImportState = &protocolsRIPDistributeListInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPDistributeListInterface{}
