/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicestunnelserverpsk code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicestunnelserverpsk

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (psk) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceStunnelServerPsk{}
	_ resource.ResourceWithConfigure   = &serviceStunnelServerPsk{}
	_ resource.ResourceWithImportState = &serviceStunnelServerPsk{}
)

// var _ resource.ResourceWithConfigValidators = &serviceStunnelServerPsk{}
// var _ resource.ResourceWithModifyPlan = &serviceStunnelServerPsk{}
// var _ resource.ResourceWithUpgradeState = &serviceStunnelServerPsk{}
// var _ resource.ResourceWithValidateConfig = &serviceStunnelServerPsk{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceStunnelServerPsk{}
