/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicestunnelserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicestunnelserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceStunnelServer{}
	_ resource.ResourceWithConfigure   = &serviceStunnelServer{}
	_ resource.ResourceWithImportState = &serviceStunnelServer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceStunnelServer{}
// var _ resource.ResourceWithModifyPlan = &serviceStunnelServer{}
// var _ resource.ResourceWithUpgradeState = &serviceStunnelServer{}
// var _ resource.ResourceWithValidateConfig = &serviceStunnelServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceStunnelServer{}
