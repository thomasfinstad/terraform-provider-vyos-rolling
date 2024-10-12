// Package namedprotocolsstaticarpinterfaceaddress code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticarpinterfaceaddress

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsStaticArpInterfaceAddress{}
	_ resource.ResourceWithConfigure = &protocolsStaticArpInterfaceAddress{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticArpInterfaceAddress{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticArpInterfaceAddress{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticArpInterfaceAddress{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticArpInterfaceAddress{}
// var _ resource.ResourceWithImportState = &protocolsStaticArpInterfaceAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticArpInterfaceAddress{}
