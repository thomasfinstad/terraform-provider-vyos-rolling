// Package namedprotocolsripinterfaceauthenticationmdfive code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsripinterfaceauthenticationmdfive

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsRIPInterfaceAuthenticationMdfive) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_protocols_rip_interface_authentication_md5"
	resp.TypeName = r.ResourceName
}