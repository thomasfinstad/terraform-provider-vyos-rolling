// Package globalvpnopenconnectaccountingmode code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectaccountingmode

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnOpenconnectAccountingMode{}
	_ resource.ResourceWithConfigure = &vpnOpenconnectAccountingMode{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectAccountingMode{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectAccountingMode{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectAccountingMode{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectAccountingMode{}
// var _ resource.ResourceWithImportState = &vpnOpenconnectAccountingMode{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectAccountingMode{}