// Package globalservicepppoeserverlog code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeserverlog

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &servicePppoeServerLog{}
	_ resource.ResourceWithConfigure = &servicePppoeServerLog{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerLog{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerLog{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerLog{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerLog{}
// var _ resource.ResourceWithImportState = &servicePppoeServerLog{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerLog{}
