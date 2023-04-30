// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfInterface describes the resource data model.
type VrfNameProtocolsOspfInterface struct {
	// LeafNodes
	VrfNameProtocolsOspfInterfaceArea               customtypes.CustomStringValue `tfsdk:"area" json:"area,omitempty"`
	VrfNameProtocolsOspfInterfaceDeadInterval       customtypes.CustomStringValue `tfsdk:"dead_interval" json:"dead-interval,omitempty"`
	VrfNameProtocolsOspfInterfaceHelloInterval      customtypes.CustomStringValue `tfsdk:"hello_interval" json:"hello-interval,omitempty"`
	VrfNameProtocolsOspfInterfaceRetransmitInterval customtypes.CustomStringValue `tfsdk:"retransmit_interval" json:"retransmit-interval,omitempty"`
	VrfNameProtocolsOspfInterfaceTransmitDelay      customtypes.CustomStringValue `tfsdk:"transmit_delay" json:"transmit-delay,omitempty"`
	VrfNameProtocolsOspfInterfaceCost               customtypes.CustomStringValue `tfsdk:"cost" json:"cost,omitempty"`
	VrfNameProtocolsOspfInterfaceMtuIgnore          customtypes.CustomStringValue `tfsdk:"mtu_ignore" json:"mtu-ignore,omitempty"`
	VrfNameProtocolsOspfInterfacePriority           customtypes.CustomStringValue `tfsdk:"priority" json:"priority,omitempty"`
	VrfNameProtocolsOspfInterfaceBandwIDth          customtypes.CustomStringValue `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
	VrfNameProtocolsOspfInterfaceHelloMultIPlier    customtypes.CustomStringValue `tfsdk:"hello_multiplier" json:"hello-multiplier,omitempty"`
	VrfNameProtocolsOspfInterfaceNetwork            customtypes.CustomStringValue `tfsdk:"network" json:"network,omitempty"`

	// TagNodes

	// Nodes
	VrfNameProtocolsOspfInterfaceAuthentication types.Object `tfsdk:"authentication" json:"authentication,omitempty"`
	VrfNameProtocolsOspfInterfaceBfd            types.Object `tfsdk:"bfd" json:"bfd,omitempty"`
	VrfNameProtocolsOspfInterfacePassive        types.Object `tfsdk:"passive" json:"passive,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfInterface) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"area": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable OSPF on this interface

|  Format  |  Description  |
|----------|---------------|
|  u32  |  OSPF area ID as decimal notation  |
|  ipv4  |  OSPF area ID in IP address notation  |

`,
		},

		"dead_interval": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interval after which a neighbor is declared dead

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Neighbor dead interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`40`),
			Computed: true,
		},

		"hello_interval": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interval between hello packets

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Hello interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"retransmit_interval": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interval between retransmitting lost link state advertisements

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Retransmit interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"transmit_delay": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Link state transmit delay

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Link state transmit delay (seconds)  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"cost": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interface cost

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  OSPF interface cost  |

`,
		},

		"mtu_ignore": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable Maximum Transmission Unit (MTU) mismatch detection

`,
		},

		"priority": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Router priority

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  OSPF router priority cost  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"bandwidth": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interface bandwidth (Mbit/s)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-100000  |  Bandwidth in Megabit/sec (for calculating OSPF cost)  |

`,
		},

		"hello_multiplier": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Hello multiplier factor

|  Format  |  Description  |
|----------|---------------|
|  u32:1-10  |  Number of Hellos to send each second  |

`,
		},

		"network": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Network type

|  Format  |  Description  |
|----------|---------------|
|  broadcast  |  Broadcast network type  |
|  non-broadcast  |  Non-broadcast network type  |
|  point-to-multipoint  |  Point-to-multipoint network type  |
|  point-to-point  |  Point-to-point network type  |

`,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfaceAuthentication{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfaceBfd{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
		},

		"passive": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfacePassive{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Suppress routing updates on an interface

`,
		},
	}
}
