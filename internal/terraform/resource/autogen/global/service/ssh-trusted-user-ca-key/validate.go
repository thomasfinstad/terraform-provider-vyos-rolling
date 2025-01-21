/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicetcptrustedusercakey code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicetcptrustedusercakey

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (trusted-user-ca-key) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceTCPTrustedUserCaKey{}
	_ resource.ResourceWithConfigure   = &serviceTCPTrustedUserCaKey{}
	_ resource.ResourceWithImportState = &serviceTCPTrustedUserCaKey{}
)

// var _ resource.ResourceWithConfigValidators = &serviceTCPTrustedUserCaKey{}
// var _ resource.ResourceWithModifyPlan = &serviceTCPTrustedUserCaKey{}
// var _ resource.ResourceWithUpgradeState = &serviceTCPTrustedUserCaKey{}
// var _ resource.ResourceWithValidateConfig = &serviceTCPTrustedUserCaKey{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceTCPTrustedUserCaKey{}
