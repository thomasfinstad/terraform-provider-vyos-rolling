// Package namedserviceipoeserverclientippool code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverclientippool

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceIPoeServerClientIPPool{}
	_ resource.ResourceWithConfigure = &serviceIPoeServerClientIPPool{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerClientIPPool{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerClientIPPool{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerClientIPPool{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerClientIPPool{}
// var _ resource.ResourceWithImportState = &serviceIPoeServerClientIPPool{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerClientIPPool{}
