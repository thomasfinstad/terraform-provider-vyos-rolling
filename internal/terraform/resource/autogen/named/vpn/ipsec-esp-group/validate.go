/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvpnipsecespgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecespgroup

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnIPsecEspGroup{}
	_ resource.ResourceWithConfigure   = &vpnIPsecEspGroup{}
	_ resource.ResourceWithImportState = &vpnIPsecEspGroup{}
)

// var _ resource.ResourceWithConfigValidators = &vpnIPsecEspGroup{}
// var _ resource.ResourceWithModifyPlan = &vpnIPsecEspGroup{}
// var _ resource.ResourceWithUpgradeState = &vpnIPsecEspGroup{}
// var _ resource.ResourceWithValidateConfig = &vpnIPsecEspGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnIPsecEspGroup{}
