// Package globalhighavailability code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalhighavailability

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &highAvailability{}
	_ resource.ResourceWithConfigure = &highAvailability{}
)

// var _ resource.ResourceWithConfigValidators = &highAvailability{}
// var _ resource.ResourceWithModifyPlan = &highAvailability{}
// var _ resource.ResourceWithUpgradeState = &highAvailability{}
// var _ resource.ResourceWithValidateConfig = &highAvailability{}
// var _ resource.ResourceWithImportState = &highAvailability{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &highAvailability{}
