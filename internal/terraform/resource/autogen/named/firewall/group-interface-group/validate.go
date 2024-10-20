/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedfirewallgroupinterfacegroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupinterfacegroup

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallGroupInterfaceGroup{}
	_ resource.ResourceWithConfigure   = &firewallGroupInterfaceGroup{}
	_ resource.ResourceWithImportState = &firewallGroupInterfaceGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupInterfaceGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupInterfaceGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupInterfaceGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupInterfaceGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupInterfaceGroup{}
