// Package namedservicelldpinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicelldpinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceLldpInterface{}
	_ resource.ResourceWithConfigure   = &serviceLldpInterface{}
	_ resource.ResourceWithImportState = &serviceLldpInterface{}
)

// var _ resource.ResourceWithConfigValidators = &serviceLldpInterface{}
// var _ resource.ResourceWithModifyPlan = &serviceLldpInterface{}
// var _ resource.ResourceWithUpgradeState = &serviceLldpInterface{}
// var _ resource.ResourceWithValidateConfig = &serviceLldpInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceLldpInterface{}
