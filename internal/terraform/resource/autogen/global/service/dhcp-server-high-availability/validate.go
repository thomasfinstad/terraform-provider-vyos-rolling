// Package globalservicedhcpserverhighavailability code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicedhcpserverhighavailability

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDhcpServerHighAvailability{}
	_ resource.ResourceWithConfigure = &serviceDhcpServerHighAvailability{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpServerHighAvailability{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpServerHighAvailability{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpServerHighAvailability{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpServerHighAvailability{}
// var _ resource.ResourceWithImportState = &serviceDhcpServerHighAvailability{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpServerHighAvailability{}
