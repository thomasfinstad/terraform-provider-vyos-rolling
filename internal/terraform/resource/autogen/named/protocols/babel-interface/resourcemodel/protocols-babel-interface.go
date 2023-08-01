// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBabelInterface describes the resource data model.
type ProtocolsBabelInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBabelInterfaceType             types.String `tfsdk:"type" vyos:"type,omitempty"`
	LeafProtocolsBabelInterfaceSplitHorizon     types.String `tfsdk:"split_horizon" vyos:"split-horizon,omitempty"`
	LeafProtocolsBabelInterfaceHelloInterval    types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafProtocolsBabelInterfaceUpdateInterval   types.Number `tfsdk:"update_interval" vyos:"update-interval,omitempty"`
	LeafProtocolsBabelInterfaceRxcost           types.Number `tfsdk:"rxcost" vyos:"rxcost,omitempty"`
	LeafProtocolsBabelInterfaceRttDecay         types.Number `tfsdk:"rtt_decay" vyos:"rtt-decay,omitempty"`
	LeafProtocolsBabelInterfaceRttMin           types.Number `tfsdk:"rtt_min" vyos:"rtt-min,omitempty"`
	LeafProtocolsBabelInterfaceRttMax           types.Number `tfsdk:"rtt_max" vyos:"rtt-max,omitempty"`
	LeafProtocolsBabelInterfaceMaxRttPenalty    types.Number `tfsdk:"max_rtt_penalty" vyos:"max-rtt-penalty,omitempty"`
	LeafProtocolsBabelInterfaceEnableTimestamps types.Bool   `tfsdk:"enable_timestamps" vyos:"enable-timestamps,omitempty"`
	LeafProtocolsBabelInterfaceChannel          types.String `tfsdk:"channel" vyos:"channel,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelInterface) GetVyosPath() []string {
	return []string{
		"protocols",

		"babel",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},

		// LeafNodes

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  auto  &emsp; |  Automatically detect interface type  |
    |  wired  &emsp; |  Wired interface  |
    |  wireless  &emsp; |  Wireless interface  |

`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		"split_horizon": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Split horizon parameters

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  default  &emsp; |  Enable on wired interfaces, and disable on wireless interfaces  |
    |  enable  &emsp; |  Enable split horizon processing  |
    |  disable  &emsp; |  Disable split horizon processing  |

`,

			// Default:          stringdefault.StaticString(`default`),
			Computed: true,
		},

		"hello_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time between scheduled hellos

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 20-655340  &emsp; |  Milliseconds  |

`,

			// Default:          stringdefault.StaticString(`4000`),
			Computed: true,
		},

		"update_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time between scheduled updates

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 20-655340  &emsp; |  Milliseconds  |

`,

			// Default:          stringdefault.StaticString(`20000`),
			Computed: true,
		},

		"rxcost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Base receive cost for this interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65534  &emsp; |  Base receive cost  |

`,
		},

		"rtt_decay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Decay factor for exponential moving average of RTT samples

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-256  &emsp; |  Decay factor, in units of 1/256  |

`,

			// Default:          stringdefault.StaticString(`42`),
			Computed: true,
		},

		"rtt_min": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum RTT

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Milliseconds  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"rtt_max": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum RTT

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Milliseconds  |

`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"max_rtt_penalty": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum additional cost due to RTT

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Milliseconds (0 to disable the use of RTT-based cost)  |

`,

			// Default:          stringdefault.StaticString(`150`),
			Computed: true,
		},

		"enable_timestamps": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable timestamps with each Hello and IHU message in order to compute RTT values

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"channel": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Channel number for diversity routing

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-254  &emsp; |  Interfaces with a channel number interfere with interfering interfaces and interfaces with the same channel number  |
    |  interfering  &emsp; |  Interfering interfaces are assumed to interfere with all other channels except non-interfering channels  |
    |  non-interfering  &emsp; |  Non-interfering interfaces only interfere with themselves  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBabelInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBabelInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
