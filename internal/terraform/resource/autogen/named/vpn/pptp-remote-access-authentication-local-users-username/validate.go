/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvpnpptpremoteaccessauthenticationlocalusersusername code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnpptpremoteaccessauthenticationlocalusersusername

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (username) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnPptpRemoteAccessAuthenticationLocalUsersUsername{}
	_ resource.ResourceWithConfigure   = &vpnPptpRemoteAccessAuthenticationLocalUsersUsername{}
	_ resource.ResourceWithImportState = &vpnPptpRemoteAccessAuthenticationLocalUsersUsername{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessAuthenticationLocalUsersUsername{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessAuthenticationLocalUsersUsername{}
