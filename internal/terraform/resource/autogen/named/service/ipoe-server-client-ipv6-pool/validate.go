// Package namedserviceipoeserverclientipvsixpool code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverclientipvsixpool

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceIPoeServerClientIPvsixPool{}
	_ resource.ResourceWithConfigure = &serviceIPoeServerClientIPvsixPool{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerClientIPvsixPool{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerClientIPvsixPool{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerClientIPvsixPool{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerClientIPvsixPool{}
// var _ resource.ResourceWithImportState = &serviceIPoeServerClientIPvsixPool{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerClientIPvsixPool{}