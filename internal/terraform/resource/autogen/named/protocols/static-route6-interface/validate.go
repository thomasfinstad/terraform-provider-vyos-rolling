// Package namedprotocolsstaticroutesixinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticroutesixinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticRoutesixInterface{}
	_ resource.ResourceWithConfigure   = &protocolsStaticRoutesixInterface{}
	_ resource.ResourceWithImportState = &protocolsStaticRoutesixInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticRoutesixInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticRoutesixInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticRoutesixInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticRoutesixInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticRoutesixInterface{}
