// Package namedservicepppoeserverclientipvsixpoolprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverclientipvsixpoolprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &servicePppoeServerClientIPvsixPoolPrefix{}
	_ resource.ResourceWithConfigure   = &servicePppoeServerClientIPvsixPoolPrefix{}
	_ resource.ResourceWithImportState = &servicePppoeServerClientIPvsixPoolPrefix{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerClientIPvsixPoolPrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerClientIPvsixPoolPrefix{}
