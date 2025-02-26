/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsripdistributelistinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsripdistributelistinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (interface) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPDistributeListInterface{}
	_ resource.ResourceWithConfigure   = &protocolsRIPDistributeListInterface{}
	_ resource.ResourceWithImportState = &protocolsRIPDistributeListInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPDistributeListInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPDistributeListInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPDistributeListInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPDistributeListInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPDistributeListInterface{}
