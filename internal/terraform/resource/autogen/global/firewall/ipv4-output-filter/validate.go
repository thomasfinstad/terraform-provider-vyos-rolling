/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalfirewallipvfouroutputfilter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvfouroutputfilter

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (filter) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallIPvfourOutputFilter{}
	_ resource.ResourceWithConfigure   = &firewallIPvfourOutputFilter{}
	_ resource.ResourceWithImportState = &firewallIPvfourOutputFilter{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourOutputFilter{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourOutputFilter{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourOutputFilter{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourOutputFilter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourOutputFilter{}
