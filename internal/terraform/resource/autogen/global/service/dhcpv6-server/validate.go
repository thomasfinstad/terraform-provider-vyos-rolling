/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicedhcpvsixserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicedhcpvsixserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (dhcpv6-server) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDhcpvsixServer{}
	_ resource.ResourceWithConfigure   = &serviceDhcpvsixServer{}
	_ resource.ResourceWithImportState = &serviceDhcpvsixServer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpvsixServer{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpvsixServer{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpvsixServer{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpvsixServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpvsixServer{}
