// Package namedprotocolsigmpproxyinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsigmpproxyinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIgmpProxyInterface{}
	_ resource.ResourceWithConfigure = &protocolsIgmpProxyInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIgmpProxyInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsIgmpProxyInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsIgmpProxyInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsIgmpProxyInterface{}
// var _ resource.ResourceWithImportState = &protocolsIgmpProxyInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIgmpProxyInterface{}
