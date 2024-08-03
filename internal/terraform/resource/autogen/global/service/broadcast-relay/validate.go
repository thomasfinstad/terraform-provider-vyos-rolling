// Package globalservicebroadcastrelay code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicebroadcastrelay

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceBroadcastRelay{}
	_ resource.ResourceWithConfigure = &serviceBroadcastRelay{}
)

// var _ resource.ResourceWithConfigValidators = &serviceBroadcastRelay{}
// var _ resource.ResourceWithModifyPlan = &serviceBroadcastRelay{}
// var _ resource.ResourceWithUpgradeState = &serviceBroadcastRelay{}
// var _ resource.ResourceWithValidateConfig = &serviceBroadcastRelay{}
// var _ resource.ResourceWithImportState = &serviceBroadcastRelay{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceBroadcastRelay{}
