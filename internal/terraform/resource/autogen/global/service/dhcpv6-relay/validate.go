// Package globalservicedhcpvsixrelay code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicedhcpvsixrelay

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDhcpvsixRelay{}
	_ resource.ResourceWithConfigure = &serviceDhcpvsixRelay{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpvsixRelay{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpvsixRelay{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpvsixRelay{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpvsixRelay{}
// var _ resource.ResourceWithImportState = &serviceDhcpvsixRelay{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpvsixRelay{}
