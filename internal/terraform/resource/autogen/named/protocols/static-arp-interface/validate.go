// Package namedprotocolsstaticarpinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticarpinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsStaticArpInterface{}
	_ resource.ResourceWithConfigure = &protocolsStaticArpInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticArpInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticArpInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticArpInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticArpInterface{}
// var _ resource.ResourceWithImportState = &protocolsStaticArpInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticArpInterface{}
