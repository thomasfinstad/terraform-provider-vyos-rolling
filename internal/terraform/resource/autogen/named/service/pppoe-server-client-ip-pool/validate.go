// Package namedservicepppoeserverclientippool code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverclientippool

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &servicePppoeServerClientIPPool{}
	_ resource.ResourceWithConfigure = &servicePppoeServerClientIPPool{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerClientIPPool{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerClientIPPool{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerClientIPPool{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerClientIPPool{}
// var _ resource.ResourceWithImportState = &servicePppoeServerClientIPPool{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerClientIPPool{}
