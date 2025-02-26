/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameprotocolsospfneighbor code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfneighbor

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (neighbor) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsOspfNeighbor{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsOspfNeighbor{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsOspfNeighbor{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfNeighbor{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfNeighbor{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfNeighbor{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfNeighbor{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfNeighbor{}
