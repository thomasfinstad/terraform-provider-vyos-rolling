// Package namedhighavailabilityvrrpgroupexcludedaddress code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedhighavailabilityvrrpgroupexcludedaddress

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &highAvailabilityVrrpGroupExcludedAddress{}
	_ resource.ResourceWithConfigure   = &highAvailabilityVrrpGroupExcludedAddress{}
	_ resource.ResourceWithImportState = &highAvailabilityVrrpGroupExcludedAddress{}
)

// var _ resource.ResourceWithConfigValidators = &highAvailabilityVrrpGroupExcludedAddress{}
// var _ resource.ResourceWithModifyPlan = &highAvailabilityVrrpGroupExcludedAddress{}
// var _ resource.ResourceWithUpgradeState = &highAvailabilityVrrpGroupExcludedAddress{}
// var _ resource.ResourceWithValidateConfig = &highAvailabilityVrrpGroupExcludedAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &highAvailabilityVrrpGroupExcludedAddress{}
