/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsisisredistributeipvfourconnectedlevelone code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourconnectedlevelone

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (level-1) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisRedistributeIPvfourConnectedLevelOne{}
	_ resource.ResourceWithConfigure   = &protocolsIsisRedistributeIPvfourConnectedLevelOne{}
	_ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvfourConnectedLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvfourConnectedLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvfourConnectedLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvfourConnectedLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvfourConnectedLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvfourConnectedLevelOne{}
