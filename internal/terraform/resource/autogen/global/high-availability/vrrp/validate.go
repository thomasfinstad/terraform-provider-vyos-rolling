// Package globalhighavailabilityvrrp code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalhighavailabilityvrrp

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &highAvailabilityVrrp{}
	_ resource.ResourceWithConfigure = &highAvailabilityVrrp{}
)

// var _ resource.ResourceWithConfigValidators = &highAvailabilityVrrp{}
// var _ resource.ResourceWithModifyPlan = &highAvailabilityVrrp{}
// var _ resource.ResourceWithUpgradeState = &highAvailabilityVrrp{}
// var _ resource.ResourceWithValidateConfig = &highAvailabilityVrrp{}
// var _ resource.ResourceWithImportState = &highAvailabilityVrrp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &highAvailabilityVrrp{}
