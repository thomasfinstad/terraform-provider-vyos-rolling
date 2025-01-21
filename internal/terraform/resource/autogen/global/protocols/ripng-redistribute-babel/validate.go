/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsripngredistributebabel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripngredistributebabel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (babel) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPngRedistributeBabel{}
	_ resource.ResourceWithConfigure   = &protocolsRIPngRedistributeBabel{}
	_ resource.ResourceWithImportState = &protocolsRIPngRedistributeBabel{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPngRedistributeBabel{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPngRedistributeBabel{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPngRedistributeBabel{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPngRedistributeBabel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPngRedistributeBabel{}
