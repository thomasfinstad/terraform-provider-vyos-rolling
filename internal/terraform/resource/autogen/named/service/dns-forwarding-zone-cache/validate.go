// Package namedservicednsforwardingzonecache code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingzonecache

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDNSForwardingZoneCache{}
	_ resource.ResourceWithConfigure   = &serviceDNSForwardingZoneCache{}
	_ resource.ResourceWithImportState = &serviceDNSForwardingZoneCache{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSForwardingZoneCache{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSForwardingZoneCache{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSForwardingZoneCache{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSForwardingZoneCache{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSForwardingZoneCache{}
