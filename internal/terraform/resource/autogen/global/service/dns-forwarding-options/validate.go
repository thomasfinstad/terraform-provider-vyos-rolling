// Package globalservicednsforwardingoptions code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicednsforwardingoptions

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDNSForwardingOptions{}
	_ resource.ResourceWithConfigure = &serviceDNSForwardingOptions{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSForwardingOptions{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSForwardingOptions{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSForwardingOptions{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSForwardingOptions{}
// var _ resource.ResourceWithImportState = &serviceDNSForwardingOptions{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSForwardingOptions{}
