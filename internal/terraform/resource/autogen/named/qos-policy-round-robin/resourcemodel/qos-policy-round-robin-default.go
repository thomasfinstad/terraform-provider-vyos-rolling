// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyRoundRobinDefault describes the resource data model.
type QosPolicyRoundRobinDefault struct {
	// LeafNodes
	QosPolicyRoundRobinDefaultCodelQuantum customtypes.CustomStringValue `tfsdk:"codel_quantum" json:"codel-quantum,omitempty"`
	QosPolicyRoundRobinDefaultFlows        customtypes.CustomStringValue `tfsdk:"flows" json:"flows,omitempty"`
	QosPolicyRoundRobinDefaultInterval     customtypes.CustomStringValue `tfsdk:"interval" json:"interval,omitempty"`
	QosPolicyRoundRobinDefaultQueueLimit   customtypes.CustomStringValue `tfsdk:"queue_limit" json:"queue-limit,omitempty"`
	QosPolicyRoundRobinDefaultQueueType    customtypes.CustomStringValue `tfsdk:"queue_type" json:"queue-type,omitempty"`
	QosPolicyRoundRobinDefaultTarget       customtypes.CustomStringValue `tfsdk:"target" json:"target,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyRoundRobinDefault) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"codel_quantum": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Deficit in the fair queuing algorithm

|  Format  |  Description  |
|----------|---------------|
|  u32:0-1048576  |  Number of bytes used as 'deficit'  |

`,

			// Default:          stringdefault.StaticString(`1514`),
			Computed: true,
		},

		"flows": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Number of flows into which the incoming packets are classified

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65536  |  Number of flows  |

`,

			// Default:          stringdefault.StaticString(`1024`),
			Computed: true,
		},

		"interval": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interval used to measure the delay

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Interval in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"queue_limit": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum queue size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Queue size in packets  |

`,
		},

		"queue_type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Queue type for default traffic

|  Format  |  Description  |
|----------|---------------|
|  drop-tail  |  First-In-First-Out (FIFO)  |
|  fair-queue  |  Stochastic Fair Queue (SFQ)  |
|  fq-codel  |  Fair Queue Codel  |
|  priority  |  Priority queuing  |
|  random-detect  |  Random Early Detection (RED)  |

`,

			// Default:          stringdefault.StaticString(`fair-queue`),
			Computed: true,
		},

		"target": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Acceptable minimum standing/persistent queue delay

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Queue delay in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
