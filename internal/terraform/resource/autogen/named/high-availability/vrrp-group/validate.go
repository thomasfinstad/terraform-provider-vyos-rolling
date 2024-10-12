// Package namedhighavailabilityvrrpgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedhighavailabilityvrrpgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &highAvailabilityVrrpGroup{}
	_ resource.ResourceWithConfigure = &highAvailabilityVrrpGroup{}
)

// var _ resource.ResourceWithConfigValidators = &highAvailabilityVrrpGroup{}
// var _ resource.ResourceWithModifyPlan = &highAvailabilityVrrpGroup{}
// var _ resource.ResourceWithUpgradeState = &highAvailabilityVrrpGroup{}
// var _ resource.ResourceWithValidateConfig = &highAvailabilityVrrpGroup{}
// var _ resource.ResourceWithImportState = &highAvailabilityVrrpGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &highAvailabilityVrrpGroup{}
