// Package namedprotocolsospfinterfaceauthenticationmdfivekeyid code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfinterfaceauthenticationmdfivekeyid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsOspfInterfaceAuthenticationMdfiveKeyID) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_ospf_interface_authentication_md5_key_id"
}
