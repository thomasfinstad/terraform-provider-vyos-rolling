// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyNetworkEmulator describes the resource data model.
type QosPolicyNetworkEmulator struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"network_emulator_id" vyos:",self-id"`

	// LeafNodes
	LeafQosPolicyNetworkEmulatorDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyNetworkEmulatorBandwIDth   types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyNetworkEmulatorDelay       types.String `tfsdk:"delay" vyos:"delay,omitempty"`
	LeafQosPolicyNetworkEmulatorCorruption  types.String `tfsdk:"corruption" vyos:"corruption,omitempty"`
	LeafQosPolicyNetworkEmulatorDuplicate   types.String `tfsdk:"duplicate" vyos:"duplicate,omitempty"`
	LeafQosPolicyNetworkEmulatorLoss        types.String `tfsdk:"loss" vyos:"loss,omitempty"`
	LeafQosPolicyNetworkEmulatorReordering  types.String `tfsdk:"reordering" vyos:"reordering,omitempty"`
	LeafQosPolicyNetworkEmulatorQueueLimit  types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o QosPolicyNetworkEmulator) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o QosPolicyNetworkEmulator) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyNetworkEmulator) GetVyosPath() []string {
	return []string{
		"qos",

		"policy",

		"network-emulator",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyNetworkEmulator) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"network_emulator_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Network emulator policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Policy name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Bits per second  |
    |  <number>bit  &emsp; |  Bits per second  |
    |  <number>kbit  &emsp; |  Kilobits per second  |
    |  <number>mbit  &emsp; |  Megabits per second  |
    |  <number>gbit  &emsp; |  Gigabits per second  |
    |  <number>tbit  &emsp; |  Terabits per second  |
    |  <number>%  &emsp; |  Percentage of interface link speed  |

`,
		},

		"delay": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adds delay to packets outgoing to chosen network interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Time in milliseconds  |

`,
		},

		"corruption": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Introducing error in a random position for chosen percent of packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Percentage of packets affected  |

`,
		},

		"duplicate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Cosen percent of packets is duplicated before queuing them

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Percentage of packets affected  |

`,
		},

		"loss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Add independent loss probability to the packets outgoing to chosen network interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Percentage of packets affected  |

`,
		},

		"reordering": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Emulated packet reordering percentage

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Percentage of packets affected  |

`,
		},

		"queue_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967295  &emsp; |  Queue size in packets  |

`,
		},

		// Nodes

	}
}
