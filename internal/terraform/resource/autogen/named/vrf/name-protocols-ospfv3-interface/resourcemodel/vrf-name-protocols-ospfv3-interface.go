// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfvthreeInterface describes the resource data model.
type VrfNameProtocolsOspfvthreeInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeInterfaceArea               types.String `tfsdk:"area" vyos:"area,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfaceDeadInterval       types.Number `tfsdk:"dead_interval" vyos:"dead-interval,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfaceHelloInterval      types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfaceRetransmitInterval types.Number `tfsdk:"retransmit_interval" vyos:"retransmit-interval,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfaceTransmitDelay      types.Number `tfsdk:"transmit_delay" vyos:"transmit-delay,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfaceCost               types.Number `tfsdk:"cost" vyos:"cost,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfaceMtuIgnore          types.Bool   `tfsdk:"mtu_ignore" vyos:"mtu-ignore,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfacePriority           types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfaceIfmtu              types.Number `tfsdk:"ifmtu" vyos:"ifmtu,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfaceInstanceID         types.Number `tfsdk:"instance_id" vyos:"instance-id,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfaceNetwork            types.String `tfsdk:"network" vyos:"network,omitempty"`
	LeafVrfNameProtocolsOspfvthreeInterfacePassive            types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfvthreeInterfaceBfd *VrfNameProtocolsOspfvthreeInterfaceBfd `tfsdk:"bfd" vyos:"bfd,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfvthreeInterface) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospfv3",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Enable routing on an IPv6 interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface used for routing information exchange  |

`,
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
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

		"ifmtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface MTU

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Interface MTU  |

`,
		},

		"instance_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Instance ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-255  &emsp; |  Instance Id  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  broadcast  &emsp; |  Broadcast network type  |
    |  point-to-point  &emsp; |  Point-to-point network type  |

`,
		},

		"passive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Configure passive mode for interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeInterfaceBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfvthreeInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfvthreeInterface) UnmarshalJSON(_ []byte) error {
	return nil
}