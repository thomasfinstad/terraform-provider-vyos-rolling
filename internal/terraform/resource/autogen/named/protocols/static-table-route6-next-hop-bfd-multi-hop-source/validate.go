/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedprotocolsstatictableroutesixnexthopbfdmultihopsource code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictableroutesixnexthopbfdmultihopsource

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticTableRoutesixNextHopBfdMultiHopSource{}
	_ resource.ResourceWithConfigure   = &protocolsStaticTableRoutesixNextHopBfdMultiHopSource{}
	_ resource.ResourceWithImportState = &protocolsStaticTableRoutesixNextHopBfdMultiHopSource{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticTableRoutesixNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticTableRoutesixNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticTableRoutesixNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticTableRoutesixNextHopBfdMultiHopSource{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticTableRoutesixNextHopBfdMultiHopSource{}
