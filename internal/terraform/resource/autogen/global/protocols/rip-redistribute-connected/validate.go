/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsripredistributeconnected code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripredistributeconnected

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (connected) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPRedistributeConnected{}
	_ resource.ResourceWithConfigure   = &protocolsRIPRedistributeConnected{}
	_ resource.ResourceWithImportState = &protocolsRIPRedistributeConnected{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPRedistributeConnected{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPRedistributeConnected{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPRedistributeConnected{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPRedistributeConnected{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPRedistributeConnected{}
