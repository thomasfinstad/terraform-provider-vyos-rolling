// Package globalprotocolsospfvthreeredistributeisis code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfvthreeredistributeisis

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfvthreeRedistributeIsis{}
	_ resource.ResourceWithConfigure   = &protocolsOspfvthreeRedistributeIsis{}
	_ resource.ResourceWithImportState = &protocolsOspfvthreeRedistributeIsis{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfvthreeRedistributeIsis{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfvthreeRedistributeIsis{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfvthreeRedistributeIsis{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfvthreeRedistributeIsis{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfvthreeRedistributeIsis{}
