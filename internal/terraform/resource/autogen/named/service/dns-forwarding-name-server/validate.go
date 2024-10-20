/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicednsforwardingnameserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingnameserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDNSForwardingNameServer{}
	_ resource.ResourceWithConfigure   = &serviceDNSForwardingNameServer{}
	_ resource.ResourceWithImportState = &serviceDNSForwardingNameServer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDNSForwardingNameServer{}
// var _ resource.ResourceWithModifyPlan = &serviceDNSForwardingNameServer{}
// var _ resource.ResourceWithUpgradeState = &serviceDNSForwardingNameServer{}
// var _ resource.ResourceWithValidateConfig = &serviceDNSForwardingNameServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDNSForwardingNameServer{}
