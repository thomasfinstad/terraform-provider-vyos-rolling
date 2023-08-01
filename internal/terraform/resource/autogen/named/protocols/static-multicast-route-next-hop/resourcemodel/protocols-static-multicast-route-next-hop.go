// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsStaticMulticastRouteNextHop describes the resource data model.
type ProtocolsStaticMulticastRouteNextHop struct {
	SelfIdentifier types.String `tfsdk:"next_hop_id" vyos:",self-id"`

	ParentIDProtocolsStaticMulticastRoute types.String `tfsdk:"route" vyos:"route,parent-id"`

	// LeafNodes
	LeafProtocolsStaticMulticastRouteNextHopDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticMulticastRouteNextHop) GetVyosPath() []string {
	return []string{
		"protocols",

		"static",

		"multicast",

		"route",
		o.ParentIDProtocolsStaticMulticastRoute.ValueString(),

		"next-hop",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticMulticastRouteNextHop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"next_hop_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Nexthop IPv4 address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Nexthop IPv4 address  |

`,
		},

		"route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Configure static unicast route into MRIB for multicast RPF lookup

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Network  |

`,
		},

		// LeafNodes

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance value for this route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Distance for this route  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsStaticMulticastRouteNextHop) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsStaticMulticastRouteNextHop) UnmarshalJSON(_ []byte) error {
	return nil
}
