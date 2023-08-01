// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyCake describes the resource data model.
type QosPolicyCake struct {
	SelfIdentifier types.String `tfsdk:"cake_id" vyos:",self-id"`

	// LeafNodes
	LeafQosPolicyCakeDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyCakeBandwIDth   types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyCakeRtt         types.Number `tfsdk:"rtt" vyos:"rtt,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeQosPolicyCakeFlowIsolation *QosPolicyCakeFlowIsolation `tfsdk:"flow_isolation" vyos:"flow-isolation,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyCake) GetVyosPath() []string {
	return []string{
		"qos",

		"policy",

		"cake",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyCake) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"cake_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Common Applications Kept Enhanced (CAKE)

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

		"rtt": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Round-Trip-Time for Active Queue Management (AQM)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-3600000  &emsp; |  RTT in ms  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		// Nodes

		"flow_isolation": schema.SingleNestedAttribute{
			Attributes: QosPolicyCakeFlowIsolation{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Flow isolation settings

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyCake) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyCake) UnmarshalJSON(_ []byte) error {
	return nil
}
