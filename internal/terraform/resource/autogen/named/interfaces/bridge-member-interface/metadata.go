// Package namedinterfacesbridgememberinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbridgememberinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r interfacesBrIDgeMemberInterface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_bridge_member_interface"
}