// Package namedvrfnameprotocolsstaticroutesixinterface code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticroutesixinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vrfNameProtocolsStaticRoutesixInterface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrf_name_protocols_static_route6_interface"
}
