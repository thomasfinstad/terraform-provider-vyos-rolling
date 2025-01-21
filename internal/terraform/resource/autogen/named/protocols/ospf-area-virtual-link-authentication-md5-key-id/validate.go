/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsospfareavirtuallinkauthenticationmdfivekeyid code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfareavirtuallinkauthenticationmdfivekeyid

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (key-id) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{}
	_ resource.ResourceWithConfigure   = &protocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{}
	_ resource.ResourceWithImportState = &protocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{}
