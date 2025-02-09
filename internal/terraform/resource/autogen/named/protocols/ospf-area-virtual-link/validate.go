/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsospfareavirtuallink code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfareavirtuallink

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (virtual-link) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfAreaVirtualLink{}
	_ resource.ResourceWithConfigure   = &protocolsOspfAreaVirtualLink{}
	_ resource.ResourceWithImportState = &protocolsOspfAreaVirtualLink{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfAreaVirtualLink{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfAreaVirtualLink{}
