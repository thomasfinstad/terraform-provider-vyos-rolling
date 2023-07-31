// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsFailoverRouteNextHop describes the resource data model.
type ProtocolsFailoverRouteNextHop struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDProtocolsFailoverRoute types.String `tfsdk:"route" vyos:"route_identifier,parent-id"`

	// LeafNodes
	LeafProtocolsFailoverRouteNextHopInterface types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafProtocolsFailoverRouteNextHopMetric    types.Number `tfsdk:"metric" vyos:"metric,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsFailoverRouteNextHopCheck *ProtocolsFailoverRouteNextHopCheck `tfsdk:"check" vyos:"check,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsFailoverRouteNextHop) GetVyosPath() []string {
	return []string{
		"protocols",

		"failover",

		"route",
		o.ParentIDProtocolsFailoverRoute.ValueString(),

		"next-hop",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsFailoverRouteNextHop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Next-hop IPv4 router address

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  Next-hop router address  |

`,
		},

		"route_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Failover IPv4 route

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  IPv4 failover route  |

`,
		},

		// LeafNodes

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway interface name

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Gateway interface name  |

`,
		},

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Route metric for this gateway

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Route metric  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		// Nodes

		"check": schema.SingleNestedAttribute{
			Attributes: ProtocolsFailoverRouteNextHopCheck{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Check target options

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsFailoverRouteNextHop) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsFailoverRouteNextHop) UnmarshalJSON(_ []byte) error {
	return nil
}
