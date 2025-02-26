/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicedhcprelayrelayoptions code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicedhcprelayrelayoptions

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (relay-options) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDhcpRelayRelayOptions{}
	_ resource.ResourceWithConfigure   = &serviceDhcpRelayRelayOptions{}
	_ resource.ResourceWithImportState = &serviceDhcpRelayRelayOptions{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpRelayRelayOptions{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpRelayRelayOptions{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpRelayRelayOptions{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpRelayRelayOptions{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpRelayRelayOptions{}
