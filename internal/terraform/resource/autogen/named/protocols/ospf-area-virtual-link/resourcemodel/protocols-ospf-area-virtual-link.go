// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsOspfAreaVirtualLink describes the resource data model.
type ProtocolsOspfAreaVirtualLink struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDProtocolsOspfArea types.String `tfsdk:"area" vyos:"area_identifier,parent-id"`

	// LeafNodes
	LeafProtocolsOspfAreaVirtualLinkDeadInterval       types.Number `tfsdk:"dead_interval" vyos:"dead-interval,omitempty"`
	LeafProtocolsOspfAreaVirtualLinkHelloInterval      types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafProtocolsOspfAreaVirtualLinkRetransmitInterval types.Number `tfsdk:"retransmit_interval" vyos:"retransmit-interval,omitempty"`
	LeafProtocolsOspfAreaVirtualLinkTransmitDelay      types.Number `tfsdk:"transmit_delay" vyos:"transmit-delay,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsOspfAreaVirtualLinkAuthentication *ProtocolsOspfAreaVirtualLinkAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfAreaVirtualLink) GetVyosPath() []string {
	return []string{
		"protocols",

		"ospf",

		"area",
		o.ParentIDProtocolsOspfArea.ValueString(),

		"virtual-link",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAreaVirtualLink) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual link

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  OSPF area in dotted decimal notation  |

`,
		},

		"area_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `OSPF area settings

    |  Format  |  Description  |
    |----------|---------------|
    |  u32  |  OSPF area number in decimal notation  |
    |  ipv4  |  OSPF area number in dotted decimal notation  |

`,
		},

		// LeafNodes

		"dead_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval after which a neighbor is declared dead

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Neighbor dead interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`40`),
			Computed: true,
		},

		"hello_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between hello packets

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Hello interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"retransmit_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between retransmitting lost link state advertisements

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Retransmit interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"transmit_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Link state transmit delay

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Link state transmit delay (seconds)  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfAreaVirtualLinkAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsOspfAreaVirtualLink) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsOspfAreaVirtualLink) UnmarshalJSON(_ []byte) error {
	return nil
}
