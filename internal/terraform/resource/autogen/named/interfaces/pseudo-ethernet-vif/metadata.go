// Package namedinterfacespseudoethernetvif code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetvif

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r interfacesPseudoEthernetVif) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_pseudo_ethernet_vif"
}
