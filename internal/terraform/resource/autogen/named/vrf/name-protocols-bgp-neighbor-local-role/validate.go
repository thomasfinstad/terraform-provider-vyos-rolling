/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfnameprotocolsbgpneighborlocalrole code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpneighborlocalrole

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsBgpNeighborLocalRole{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsBgpNeighborLocalRole{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsBgpNeighborLocalRole{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpNeighborLocalRole{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpNeighborLocalRole{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpNeighborLocalRole{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpNeighborLocalRole{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpNeighborLocalRole{}
