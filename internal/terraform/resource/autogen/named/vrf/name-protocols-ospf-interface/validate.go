/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameprotocolsospfinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (interface) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsOspfInterface{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsOspfInterface{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsOspfInterface{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfInterface{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfInterface{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfInterface{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfInterface{}
