// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfInterface describes the resource data model.
type VrfNameProtocolsOspfInterface struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsOspfInterfaceArea               types.String `tfsdk:"area" vyos:"area,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceDeadInterval       types.Number `tfsdk:"dead_interval" vyos:"dead-interval,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceHelloInterval      types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceRetransmitInterval types.Number `tfsdk:"retransmit_interval" vyos:"retransmit-interval,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceTransmitDelay      types.Number `tfsdk:"transmit_delay" vyos:"transmit-delay,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceCost               types.Number `tfsdk:"cost" vyos:"cost,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceMtuIgnore          types.Bool   `tfsdk:"mtu_ignore" vyos:"mtu-ignore,omitempty"`
	LeafVrfNameProtocolsOspfInterfacePriority           types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceBandwIDth          types.Number `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceHelloMultIPlier    types.Number `tfsdk:"hello_multiplier" vyos:"hello-multiplier,omitempty"`
	LeafVrfNameProtocolsOspfInterfaceNetwork            types.String `tfsdk:"network" vyos:"network,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfInterfaceAuthentication *VrfNameProtocolsOspfInterfaceAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`
	NodeVrfNameProtocolsOspfInterfaceBfd            *VrfNameProtocolsOspfInterfaceBfd            `tfsdk:"bfd" vyos:"bfd,omitempty"`
	NodeVrfNameProtocolsOspfInterfacePassive        *VrfNameProtocolsOspfInterfacePassive        `tfsdk:"passive" vyos:"passive,omitempty"`
}

// GetID returns the resource ID
func (o VrfNameProtocolsOspfInterface) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o VrfNameProtocolsOspfInterface) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfInterface) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospf",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"area": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable OSPF on this interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  OSPF area ID as decimal notation  |
    |  ipv4  &emsp; |  OSPF area ID in IP address notation  |

`,
		},

		"dead_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval after which a neighbor is declared dead

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Neighbor dead interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`40`),
			Computed: true,
		},

		"hello_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between hello packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Hello interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"retransmit_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between retransmitting lost link state advertisements

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Retransmit interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"transmit_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Link state transmit delay

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Link state transmit delay (seconds)  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"cost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface cost

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  OSPF interface cost  |

`,
		},

		"mtu_ignore": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable Maximum Transmission Unit (MTU) mismatch detection

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Router priority

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-255  &emsp; |  OSPF router priority cost  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"bandwidth": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface bandwidth (Mbit/s)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-100000  &emsp; |  Bandwidth in Megabit/sec (for calculating OSPF cost)  |

`,
		},

		"hello_multiplier": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello multiplier factor

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-10  &emsp; |  Number of Hellos to send each second  |

`,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  broadcast  &emsp; |  Broadcast network type  |
    |  non-broadcast  &emsp; |  Non-broadcast network type  |
    |  point-to-multipoint  &emsp; |  Point-to-multipoint network type  |
    |  point-to-point  &emsp; |  Point-to-point network type  |

`,
		},

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfaceAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfaceBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
		},

		"passive": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfacePassive{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Suppress routing updates on an interface

`,
		},
	}
}
