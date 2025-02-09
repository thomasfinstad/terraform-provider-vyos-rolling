/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsstaticmroutenexthop code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticmroutenexthop

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (next-hop) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticMrouteNextHop{}
	_ resource.ResourceWithConfigure   = &protocolsStaticMrouteNextHop{}
	_ resource.ResourceWithImportState = &protocolsStaticMrouteNextHop{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticMrouteNextHop{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticMrouteNextHop{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticMrouteNextHop{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticMrouteNextHop{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticMrouteNextHop{}
