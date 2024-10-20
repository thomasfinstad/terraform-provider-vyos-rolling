/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedserviceeventhandlereventscriptenvironment code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceeventhandlereventscriptenvironment

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceEventHandlerEventScrIPtEnvironment{}
	_ resource.ResourceWithConfigure   = &serviceEventHandlerEventScrIPtEnvironment{}
	_ resource.ResourceWithImportState = &serviceEventHandlerEventScrIPtEnvironment{}
)

// var _ resource.ResourceWithConfigValidators = &serviceEventHandlerEventScrIPtEnvironment{}
// var _ resource.ResourceWithModifyPlan = &serviceEventHandlerEventScrIPtEnvironment{}
// var _ resource.ResourceWithUpgradeState = &serviceEventHandlerEventScrIPtEnvironment{}
// var _ resource.ResourceWithValidateConfig = &serviceEventHandlerEventScrIPtEnvironment{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceEventHandlerEventScrIPtEnvironment{}
