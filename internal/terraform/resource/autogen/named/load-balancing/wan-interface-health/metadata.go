// Package namedloadbalancingwaninterfacehealth code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingwaninterfacehealth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r loadBalancingWanInterfaceHealth) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_load_balancing_wan_interface_health"
}