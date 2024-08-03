// Package globalservicedhcprelay code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicedhcprelay

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDhcpRelay{}
	_ resource.ResourceWithConfigure = &serviceDhcpRelay{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpRelay{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpRelay{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpRelay{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpRelay{}
// var _ resource.ResourceWithImportState = &serviceDhcpRelay{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpRelay{}
