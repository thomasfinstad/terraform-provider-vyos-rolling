// Package globalserviceipoeserversnmp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceipoeserversnmp

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIPoeServerSnmp{}
	_ resource.ResourceWithConfigure   = &serviceIPoeServerSnmp{}
	_ resource.ResourceWithImportState = &serviceIPoeServerSnmp{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerSnmp{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerSnmp{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerSnmp{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerSnmp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerSnmp{}
