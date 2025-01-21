/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsstaticrouteinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticrouteinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (interface) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticRouteInterface{}
	_ resource.ResourceWithConfigure   = &protocolsStaticRouteInterface{}
	_ resource.ResourceWithImportState = &protocolsStaticRouteInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticRouteInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticRouteInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticRouteInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticRouteInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticRouteInterface{}
