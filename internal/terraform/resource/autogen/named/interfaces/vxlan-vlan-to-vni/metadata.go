// Package namedinterfacesvxlanvlantovni code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvxlanvlantovni

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r interfacesVxlanVlanToVni) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_vxlan_vlan_to_vni"
}
