// Package namedvpnipsecikegroup code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecikegroup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnIPsecIkeGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ipsec_ike_group"
}