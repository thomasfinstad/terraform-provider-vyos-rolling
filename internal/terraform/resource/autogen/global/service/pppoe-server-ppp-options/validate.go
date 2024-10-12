// Package globalservicepppoeserverpppoptions code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeserverpppoptions

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &servicePppoeServerPppOptions{}
	_ resource.ResourceWithConfigure = &servicePppoeServerPppOptions{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerPppOptions{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerPppOptions{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerPppOptions{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerPppOptions{}
// var _ resource.ResourceWithImportState = &servicePppoeServerPppOptions{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerPppOptions{}
