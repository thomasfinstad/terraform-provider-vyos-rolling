/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolspimrpaddress code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolspimrpaddress

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (address) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsPimRpAddress{}
	_ resource.ResourceWithConfigure   = &protocolsPimRpAddress{}
	_ resource.ResourceWithImportState = &protocolsPimRpAddress{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsPimRpAddress{}
// var _ resource.ResourceWithModifyPlan = &protocolsPimRpAddress{}
// var _ resource.ResourceWithUpgradeState = &protocolsPimRpAddress{}
// var _ resource.ResourceWithValidateConfig = &protocolsPimRpAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsPimRpAddress{}
