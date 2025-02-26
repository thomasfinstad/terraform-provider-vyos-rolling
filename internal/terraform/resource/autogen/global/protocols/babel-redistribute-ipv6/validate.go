/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbabelredistributeipvsix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbabelredistributeipvsix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (ipv6) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBabelRedistributeIPvsix{}
	_ resource.ResourceWithConfigure   = &protocolsBabelRedistributeIPvsix{}
	_ resource.ResourceWithImportState = &protocolsBabelRedistributeIPvsix{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBabelRedistributeIPvsix{}
// var _ resource.ResourceWithModifyPlan = &protocolsBabelRedistributeIPvsix{}
// var _ resource.ResourceWithUpgradeState = &protocolsBabelRedistributeIPvsix{}
// var _ resource.ResourceWithValidateConfig = &protocolsBabelRedistributeIPvsix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBabelRedistributeIPvsix{}
