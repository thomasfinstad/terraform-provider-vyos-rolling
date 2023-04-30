// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDhcpServerSharedNetworkNameSubnetRange describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnetRange struct {
	// LeafNodes
	ServiceDhcpServerSharedNetworkNameSubnetRangeStart customtypes.CustomStringValue `tfsdk:"start" json:"start,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetRangeStop  customtypes.CustomStringValue `tfsdk:"stop" json:"stop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetRange) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"start": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `First IP address for DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 start address of pool  |

`,
		},

		"stop": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Last IP address for DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 end address of pool  |

`,
		},

		// TagNodes

		// Nodes

	}
}
