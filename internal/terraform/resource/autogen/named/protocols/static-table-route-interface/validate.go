/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsstatictablerouteinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictablerouteinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (interface) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticTableRouteInterface{}
	_ resource.ResourceWithConfigure   = &protocolsStaticTableRouteInterface{}
	_ resource.ResourceWithImportState = &protocolsStaticTableRouteInterface{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticTableRouteInterface{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticTableRouteInterface{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticTableRouteInterface{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticTableRouteInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticTableRouteInterface{}
