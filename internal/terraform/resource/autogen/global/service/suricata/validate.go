// Package globalservicesuricata code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicesuricata

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceSURIcata{}
	_ resource.ResourceWithConfigure = &serviceSURIcata{}
)

// var _ resource.ResourceWithConfigValidators = &serviceSURIcata{}
// var _ resource.ResourceWithModifyPlan = &serviceSURIcata{}
// var _ resource.ResourceWithUpgradeState = &serviceSURIcata{}
// var _ resource.ResourceWithValidateConfig = &serviceSURIcata{}
// var _ resource.ResourceWithImportState = &serviceSURIcata{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceSURIcata{}
