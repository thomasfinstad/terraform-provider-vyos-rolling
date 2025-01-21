/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicepppoeserverpadodelay code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverpadodelay

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (pado-delay) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &servicePppoeServerPadoDelay{}
	_ resource.ResourceWithConfigure   = &servicePppoeServerPadoDelay{}
	_ resource.ResourceWithImportState = &servicePppoeServerPadoDelay{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerPadoDelay{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerPadoDelay{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerPadoDelay{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerPadoDelay{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerPadoDelay{}
