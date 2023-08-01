// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyFairQueue describes the resource data model.
type QosPolicyFairQueue struct {
	SelfIdentifier types.String `tfsdk:"fair_queue_id" vyos:",self-id"`

	// LeafNodes
	LeafQosPolicyFairQueueDescrIPtion  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyFairQueueHashInterval types.Number `tfsdk:"hash_interval" vyos:"hash-interval,omitempty"`
	LeafQosPolicyFairQueueQueueLimit   types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyFairQueue) GetVyosPath() []string {
	return []string{
		"qos",

		"policy",

		"fair-queue",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyFairQueue) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `fair_queue_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"fair_queue_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Stochastic Fairness Queueing

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

		"hash_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval in seconds for queue algorithm perturbation

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0  &emsp; |  No perturbation  |
    |  number: 1-127  &emsp; |  Interval in seconds for queue algorithm perturbation (advised: 10)  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"queue_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Upper limit of the SFQ

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-127  &emsp; |  Queue size in packets  |

`,

			// Default:          stringdefault.StaticString(`127`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyFairQueue) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyFairQueue) UnmarshalJSON(_ []byte) error {
	return nil
}
