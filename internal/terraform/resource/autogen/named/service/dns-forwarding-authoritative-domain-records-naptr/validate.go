/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicednsforwardingauthoritativedomainrecordsnaptr code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordsnaptr

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (naptr) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
	_ resource.ResourceWithConfigure   = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
	_ resource.ResourceWithImportState = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSForwardingAuthoritativeDomainRecordsNaptr{}
