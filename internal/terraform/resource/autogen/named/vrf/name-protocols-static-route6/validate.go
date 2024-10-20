/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfnameprotocolsstaticroutesix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticroutesix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsStaticRoutesix{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsStaticRoutesix{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsStaticRoutesix{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsStaticRoutesix{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsStaticRoutesix{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsStaticRoutesix{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsStaticRoutesix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsStaticRoutesix{}
