// Package globalprotocolsbgpparametersdampening code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpparametersdampening

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpParametersDampening{}
	_ resource.ResourceWithConfigure = &protocolsBgpParametersDampening{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpParametersDampening{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpParametersDampening{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpParametersDampening{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpParametersDampening{}
// var _ resource.ResourceWithImportState = &protocolsBgpParametersDampening{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpParametersDampening{}
