/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsfailoverroutenexthop code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsfailoverroutenexthop

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (next-hop) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsFailoverRouteNextHop{}
	_ resource.ResourceWithConfigure   = &protocolsFailoverRouteNextHop{}
	_ resource.ResourceWithImportState = &protocolsFailoverRouteNextHop{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsFailoverRouteNextHop{}
// var _ resource.ResourceWithModifyPlan = &protocolsFailoverRouteNextHop{}
// var _ resource.ResourceWithUpgradeState = &protocolsFailoverRouteNextHop{}
// var _ resource.ResourceWithValidateConfig = &protocolsFailoverRouteNextHop{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsFailoverRouteNextHop{}
