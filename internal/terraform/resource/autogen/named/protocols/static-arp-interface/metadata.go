// Package namedprotocolsstaticarpinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticarpinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsStaticArpInterface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_static_arp_interface"
}