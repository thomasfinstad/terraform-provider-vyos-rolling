// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsStaticRoutesix describes the resource data model.
type ProtocolsStaticRoutesix struct {
	// LeafNodes
	ProtocolsStaticRoutesixDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`

	// TagNodes
	ProtocolsStaticRoutesixInterface types.Map `tfsdk:"interface" json:"interface,omitempty"`
	ProtocolsStaticRoutesixNextHop   types.Map `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// Nodes
	ProtocolsStaticRoutesixBlackhole types.Object `tfsdk:"blackhole" json:"blackhole,omitempty"`
	ProtocolsStaticRoutesixReject    types.Object `tfsdk:"reject" json:"reject,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsStaticRoutesix) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		// TagNodes

		"interface": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsStaticRoutesixInterface{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `IPv6 gateway interface name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Gateway interface name  |

`,
		},

		"next_hop": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsStaticRoutesixNextHop{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `IPv6 gateway address

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Next-hop IPv6 router  |

`,
		},

		// Nodes

		"blackhole": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticRoutesixBlackhole{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Silently discard pkts when matched

`,
		},

		"reject": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticRoutesixReject{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Emit an ICMP unreachable when matched

`,
		},
	}
}
