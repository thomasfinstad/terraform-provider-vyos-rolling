// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsStaticTableRoute describes the resource data model.
type ProtocolsStaticTableRoute struct {
	SelfIdentifier types.String `tfsdk:"route_id" vyos:",self-id"`

	ParentIDProtocolsStaticTable types.String `tfsdk:"table" vyos:"table,parent-id"`

	// LeafNodes
	LeafProtocolsStaticTableRouteDhcpInterface types.String `tfsdk:"dhcp_interface" vyos:"dhcp-interface,omitempty"`
	LeafProtocolsStaticTableRouteDescrIPtion   types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsStaticTableRouteInterface bool `tfsdk:"interface" vyos:"interface,child"`
	ExistsTagProtocolsStaticTableRouteNextHop   bool `tfsdk:"next_hop" vyos:"next-hop,child"`

	// Nodes
	NodeProtocolsStaticTableRouteBlackhole *ProtocolsStaticTableRouteBlackhole `tfsdk:"blackhole" vyos:"blackhole,omitempty"`
	NodeProtocolsStaticTableRouteReject    *ProtocolsStaticTableRouteReject    `tfsdk:"reject" vyos:"reject,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticTableRoute) GetVyosPath() []string {
	return []string{
		"protocols",

		"static",

		"table",
		o.ParentIDProtocolsStaticTable.ValueString(),

		"route",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticTableRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv4 route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 static route  |

`,
		},

		"table_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Policy route table number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-200  &emsp; |  Policy route table number  |

`,
		},

		// LeafNodes

		"dhcp_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP interface supplying next-hop IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  DHCP interface name  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		// Nodes

		"blackhole": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticTableRouteBlackhole{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Silently discard pkts when matched

`,
		},

		"reject": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticTableRouteReject{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Emit an ICMP unreachable when matched

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsStaticTableRoute) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsStaticTableRoute) UnmarshalJSON(_ []byte) error {
	return nil
}
