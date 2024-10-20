/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicestunnelclient code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicestunnelclient

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceStunnelClient{}
	_ resource.ResourceWithConfigure   = &serviceStunnelClient{}
	_ resource.ResourceWithImportState = &serviceStunnelClient{}
)

// var _ resource.ResourceWithConfigValidators = &serviceStunnelClient{}
// var _ resource.ResourceWithModifyPlan = &serviceStunnelClient{}
// var _ resource.ResourceWithUpgradeState = &serviceStunnelClient{}
// var _ resource.ResourceWithValidateConfig = &serviceStunnelClient{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceStunnelClient{}
