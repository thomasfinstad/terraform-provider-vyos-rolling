/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameprotocolsospfinterfaceauthenticationmdfivekeyid code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfinterfaceauthenticationmdfivekeyid

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (key-id) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID{}
