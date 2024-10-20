/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfnameprotocolsospfareavirtuallink code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfareavirtuallink

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsOspfAreaVirtualLink{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsOspfAreaVirtualLink{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsOspfAreaVirtualLink{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfAreaVirtualLink{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfAreaVirtualLink{}
