// Package namedservicednsforwardingauthoritativedomainrecordsa code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordsa

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDNSForwardingAuthoritativeDomainRecordsA{}
	_ resource.ResourceWithConfigure = &serviceDNSForwardingAuthoritativeDomainRecordsA{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSForwardingAuthoritativeDomainRecordsA{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSForwardingAuthoritativeDomainRecordsA{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSForwardingAuthoritativeDomainRecordsA{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSForwardingAuthoritativeDomainRecordsA{}
// var _ resource.ResourceWithImportState = &serviceDNSForwardingAuthoritativeDomainRecordsA{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSForwardingAuthoritativeDomainRecordsA{}
