// Package globalservicetcprekey code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicetcprekey

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceTCPRekey{}
	_ resource.ResourceWithConfigure = &serviceTCPRekey{}
)

// var _ resource.ResourceWithConfigValidators = &serviceTCPRekey{}
// var _ resource.ResourceWithModifyPlan = &serviceTCPRekey{}
// var _ resource.ResourceWithUpgradeState = &serviceTCPRekey{}
// var _ resource.ResourceWithValidateConfig = &serviceTCPRekey{}
// var _ resource.ResourceWithImportState = &serviceTCPRekey{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceTCPRekey{}
