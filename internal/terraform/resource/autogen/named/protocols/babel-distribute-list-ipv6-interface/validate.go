// Package namedprotocolsbabeldistributelistipvsixinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbabeldistributelistipvsixinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBabelDistributeListIPvsixInterface{}
	_ resource.ResourceWithConfigure = &protocolsBabelDistributeListIPvsixInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBabelDistributeListIPvsixInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsBabelDistributeListIPvsixInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsBabelDistributeListIPvsixInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsBabelDistributeListIPvsixInterface{}
// var _ resource.ResourceWithImportState = &protocolsBabelDistributeListIPvsixInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBabelDistributeListIPvsixInterface{}
