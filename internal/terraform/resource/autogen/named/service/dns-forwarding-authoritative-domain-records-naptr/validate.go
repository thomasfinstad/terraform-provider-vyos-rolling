// Package namedservicednsforwardingauthoritativedomainrecordsnaptr code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordsnaptr

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
	_ resource.ResourceWithConfigure = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
// var _ resource.ResourceWithImportState = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
