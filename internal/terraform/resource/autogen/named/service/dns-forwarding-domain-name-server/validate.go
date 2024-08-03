// Package namedservicednsforwardingdomainnameserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingdomainnameserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDNSForwardingDomainNameServer{}
	_ resource.ResourceWithConfigure = &serviceDNSForwardingDomainNameServer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSForwardingDomainNameServer{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSForwardingDomainNameServer{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSForwardingDomainNameServer{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSForwardingDomainNameServer{}
// var _ resource.ResourceWithImportState = &serviceDNSForwardingDomainNameServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSForwardingDomainNameServer{}
