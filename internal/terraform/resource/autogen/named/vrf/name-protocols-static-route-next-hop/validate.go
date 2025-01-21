/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameprotocolsstaticroutenexthop code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticroutenexthop

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (next-hop) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsStaticRouteNextHop{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsStaticRouteNextHop{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsStaticRouteNextHop{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsStaticRouteNextHop{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsStaticRouteNextHop{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsStaticRouteNextHop{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsStaticRouteNextHop{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsStaticRouteNextHop{}
