// Package namedprotocolsstaticmulticastinterfaceroutenexthopinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticmulticastinterfaceroutenexthopinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsStaticMulticastInterfaceRouteNextHopInterface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_static_multicast_interface_route_next_hop_interface"
}
