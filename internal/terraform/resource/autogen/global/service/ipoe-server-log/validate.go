// Package globalserviceipoeserverlog code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceipoeserverlog

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceIPoeServerLog{}
	_ resource.ResourceWithConfigure = &serviceIPoeServerLog{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerLog{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerLog{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerLog{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerLog{}
// var _ resource.ResourceWithImportState = &serviceIPoeServerLog{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerLog{}
