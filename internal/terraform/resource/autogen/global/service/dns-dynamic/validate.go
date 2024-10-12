// Package globalservicednsdynamic code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicednsdynamic

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDNSDynamic{}
	_ resource.ResourceWithConfigure = &serviceDNSDynamic{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSDynamic{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSDynamic{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSDynamic{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSDynamic{}
// var _ resource.ResourceWithImportState = &serviceDNSDynamic{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSDynamic{}
