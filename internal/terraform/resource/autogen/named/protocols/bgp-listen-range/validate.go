/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedprotocolsbgplistenrange code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgplistenrange

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpListenRange{}
	_ resource.ResourceWithConfigure   = &protocolsBgpListenRange{}
	_ resource.ResourceWithImportState = &protocolsBgpListenRange{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpListenRange{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpListenRange{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpListenRange{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpListenRange{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpListenRange{}
