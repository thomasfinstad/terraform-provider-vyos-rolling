/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicednsforwardingauthoritativedomain code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomain

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (authoritative-domain) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDNSForwardingAuthoritativeDomain{}
	_ resource.ResourceWithConfigure   = &serviceDNSForwardingAuthoritativeDomain{}
	_ resource.ResourceWithImportState = &serviceDNSForwardingAuthoritativeDomain{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSForwardingAuthoritativeDomain{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSForwardingAuthoritativeDomain{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSForwardingAuthoritativeDomain{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSForwardingAuthoritativeDomain{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSForwardingAuthoritativeDomain{}
