// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyRoundRobin describes the resource data model.
type QosPolicyRoundRobin struct {
	SelfIdentifier types.String `tfsdk:"round_robin_id" vyos:",self-id"`

	// LeafNodes
	LeafQosPolicyRoundRobinDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagQosPolicyRoundRobinClass bool `tfsdk:"class" vyos:"class,child"`

	// Nodes
	NodeQosPolicyRoundRobinDefault *QosPolicyRoundRobinDefault `tfsdk:"default" vyos:"default,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyRoundRobin) GetVyosPath() []string {
	return []string{
		"qos",

		"policy",

		"round-robin",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobin) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"round_robin_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Deficit Round Robin Scheduler

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Policy name  |

`,
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

		// Nodes

		"default": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinDefault{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Default policy

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyRoundRobin) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyRoundRobin) UnmarshalJSON(_ []byte) error {
	return nil
}
