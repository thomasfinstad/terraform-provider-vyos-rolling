// Package namedserviceeventhandlerevent code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceeventhandlerevent

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceEventHandlerEvent{}
	_ resource.ResourceWithConfigure = &serviceEventHandlerEvent{}
)

// var _ resource.ResourceWithConfigValidators = &serviceEventHandlerEvent{}
// var _ resource.ResourceWithModifyPlan = &serviceEventHandlerEvent{}
// var _ resource.ResourceWithUpgradeState = &serviceEventHandlerEvent{}
// var _ resource.ResourceWithValidateConfig = &serviceEventHandlerEvent{}
// var _ resource.ResourceWithImportState = &serviceEventHandlerEvent{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceEventHandlerEvent{}