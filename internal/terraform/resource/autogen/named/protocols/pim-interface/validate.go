// Package namedprotocolspiminterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolspiminterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsPimInterface{}
	_ resource.ResourceWithConfigure = &protocolsPimInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsPimInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsPimInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsPimInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsPimInterface{}
// var _ resource.ResourceWithImportState = &protocolsPimInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsPimInterface{}
