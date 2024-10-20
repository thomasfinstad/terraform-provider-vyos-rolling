/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfnameprotocolsisisinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsisisinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsIsisInterface{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsIsisInterface{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsIsisInterface{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsIsisInterface{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsIsisInterface{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsIsisInterface{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsIsisInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsIsisInterface{}
