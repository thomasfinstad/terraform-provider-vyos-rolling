// Package globalvpnsstpauthenticationradiusdynamicauthor code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstpauthenticationradiusdynamicauthor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnSstpAuthenticationRadiusDynamicAuthor) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_sstp_authentication_radius_dynamic_author"
}
