/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsmplsldpallocationipvfour code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsldpallocationipvfour

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsMplsLdpAllocationIPvfour{}
	_ resource.ResourceWithConfigure   = &protocolsMplsLdpAllocationIPvfour{}
	_ resource.ResourceWithImportState = &protocolsMplsLdpAllocationIPvfour{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsMplsLdpAllocationIPvfour{}
// var _ resource.ResourceWithModifyPlan = &protocolsMplsLdpAllocationIPvfour{}
// var _ resource.ResourceWithUpgradeState = &protocolsMplsLdpAllocationIPvfour{}
// var _ resource.ResourceWithValidateConfig = &protocolsMplsLdpAllocationIPvfour{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsMplsLdpAllocationIPvfour{}
