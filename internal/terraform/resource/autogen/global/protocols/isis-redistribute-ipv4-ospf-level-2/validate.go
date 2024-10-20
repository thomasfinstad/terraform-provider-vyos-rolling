/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsisisredistributeipvfourospfleveltwo code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourospfleveltwo

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisRedistributeIPvfourOspfLevelTwo{}
	_ resource.ResourceWithConfigure   = &protocolsIsisRedistributeIPvfourOspfLevelTwo{}
	_ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvfourOspfLevelTwo{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvfourOspfLevelTwo{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvfourOspfLevelTwo{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvfourOspfLevelTwo{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvfourOspfLevelTwo{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvfourOspfLevelTwo{}
