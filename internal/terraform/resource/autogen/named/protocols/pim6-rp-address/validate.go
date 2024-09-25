// Package namedprotocolspimsixrpaddress code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolspimsixrpaddress

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsPimsixRpAddress{}
	_ resource.ResourceWithConfigure = &protocolsPimsixRpAddress{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsPimsixRpAddress{}
// var _ resource.ResourceWithModifyPlan = &protocolsPimsixRpAddress{}
// var _ resource.ResourceWithUpgradeState = &protocolsPimsixRpAddress{}
// var _ resource.ResourceWithValidateConfig = &protocolsPimsixRpAddress{}
// var _ resource.ResourceWithImportState = &protocolsPimsixRpAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsPimsixRpAddress{}
