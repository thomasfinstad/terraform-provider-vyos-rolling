// Package namedprotocolsripngdistributelistinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsripngdistributelistinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPngDistributeListInterface{}
	_ resource.ResourceWithConfigure = &protocolsRIPngDistributeListInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPngDistributeListInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPngDistributeListInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPngDistributeListInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPngDistributeListInterface{}
// var _ resource.ResourceWithImportState = &protocolsRIPngDistributeListInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPngDistributeListInterface{}