// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyPriorityQueueDefault describes the resource data model.
type QosPolicyPriorityQueueDefault struct {
	// LeafNodes
	LeafQosPolicyPriorityQueueDefaultCodelQuantum types.Number `tfsdk:"codel_quantum" vyos:"codel-quantum,omitempty"`
	LeafQosPolicyPriorityQueueDefaultFlows        types.Number `tfsdk:"flows" vyos:"flows,omitempty"`
	LeafQosPolicyPriorityQueueDefaultInterval     types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafQosPolicyPriorityQueueDefaultQueueLimit   types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`
	LeafQosPolicyPriorityQueueDefaultQueueType    types.String `tfsdk:"queue_type" vyos:"queue-type,omitempty"`
	LeafQosPolicyPriorityQueueDefaultTarget       types.Number `tfsdk:"target" vyos:"target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyPriorityQueueDefault) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"codel_quantum": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Deficit in the fair queuing algorithm

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-1048576  |  Number of bytes used as 'deficit'  |

`,

			// Default:          stringdefault.StaticString(`1514`),
			Computed: true,
		},

		"flows": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of flows into which the incoming packets are classified

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65536  |  Number of flows  |

`,

			// Default:          stringdefault.StaticString(`1024`),
			Computed: true,
		},

		"interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval used to measure the delay

    |  Format  |  Description  |
    |----------|---------------|
    |  u32  |  Interval in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"queue_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-4294967295  |  Queue size in packets  |

`,
		},

		"queue_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Queue type for default traffic

    |  Format  |  Description  |
    |----------|---------------|
    |  drop-tail  |  First-In-First-Out (FIFO)  |
    |  fair-queue  |  Stochastic Fair Queue (SFQ)  |
    |  fq-codel  |  Fair Queue Codel  |
    |  priority  |  Priority queuing  |
    |  random-detect  |  Random Early Detection (RED)  |

`,

			// Default:          stringdefault.StaticString(`drop-tail`),
			Computed: true,
		},

		"target": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Acceptable minimum standing/persistent queue delay

    |  Format  |  Description  |
    |----------|---------------|
    |  u32  |  Queue delay in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyPriorityQueueDefault) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyPriorityQueueDefault) UnmarshalJSON(_ []byte) error {
	return nil
}
