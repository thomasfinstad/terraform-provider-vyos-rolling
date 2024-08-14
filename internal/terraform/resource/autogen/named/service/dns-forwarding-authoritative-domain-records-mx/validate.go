// Package namedservicednsforwardingauthoritativedomainrecordsmx code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordsmx

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDNSForwardingAuthoritativeDomainRecordsMx{}
	_ resource.ResourceWithConfigure = &serviceDNSForwardingAuthoritativeDomainRecordsMx{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSForwardingAuthoritativeDomainRecordsMx{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSForwardingAuthoritativeDomainRecordsMx{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSForwardingAuthoritativeDomainRecordsMx{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSForwardingAuthoritativeDomainRecordsMx{}
// var _ resource.ResourceWithImportState = &serviceDNSForwardingAuthoritativeDomainRecordsMx{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSForwardingAuthoritativeDomainRecordsMx{}