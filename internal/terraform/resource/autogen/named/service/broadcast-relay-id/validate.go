/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicebroadcastrelayid code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicebroadcastrelayid

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (id) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceBroadcastRelayID{}
	_ resource.ResourceWithConfigure   = &serviceBroadcastRelayID{}
	_ resource.ResourceWithImportState = &serviceBroadcastRelayID{}
)

// var _ resource.ResourceWithConfigValidators = &serviceBroadcastRelayID{}
// var _ resource.ResourceWithModifyPlan = &serviceBroadcastRelayID{}
// var _ resource.ResourceWithUpgradeState = &serviceBroadcastRelayID{}
// var _ resource.ResourceWithValidateConfig = &serviceBroadcastRelayID{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceBroadcastRelayID{}
