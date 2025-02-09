/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsisisredistributeipvfourbabellevelone code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourbabellevelone

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (level-1) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
	_ resource.ResourceWithConfigure   = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
	_ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvfourBabelLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
