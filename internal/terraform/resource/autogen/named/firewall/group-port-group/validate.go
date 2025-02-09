/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedfirewallgroupportgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupportgroup

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (port-group) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallGroupPortGroup{}
	_ resource.ResourceWithConfigure   = &firewallGroupPortGroup{}
	_ resource.ResourceWithImportState = &firewallGroupPortGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupPortGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupPortGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupPortGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupPortGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupPortGroup{}
