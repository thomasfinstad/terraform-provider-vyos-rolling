/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedfirewallzonefrom code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallzonefrom

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallZoneFrom{}
	_ resource.ResourceWithConfigure   = &firewallZoneFrom{}
	_ resource.ResourceWithImportState = &firewallZoneFrom{}
)

// var _ resource.ResourceWithConfigValidators = &firewallZoneFrom{}
// var _ resource.ResourceWithModifyPlan = &firewallZoneFrom{}
// var _ resource.ResourceWithUpgradeState = &firewallZoneFrom{}
// var _ resource.ResourceWithValidateConfig = &firewallZoneFrom{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallZoneFrom{}
