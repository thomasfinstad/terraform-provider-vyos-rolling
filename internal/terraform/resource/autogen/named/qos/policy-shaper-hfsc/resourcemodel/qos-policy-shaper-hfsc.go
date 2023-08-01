// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyShaperHfsc describes the resource data model.
type QosPolicyShaperHfsc struct {
	SelfIdentifier types.String `tfsdk:"shaper_hfsc_id" vyos:",self-id"`

	// LeafNodes
	LeafQosPolicyShaperHfscDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyShaperHfscBandwIDth   types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagQosPolicyShaperHfscClass bool `tfsdk:"class" vyos:"class,child"`

	// Nodes
	NodeQosPolicyShaperHfscDefault *QosPolicyShaperHfscDefault `tfsdk:"default" vyos:"default,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyShaperHfsc) GetVyosPath() []string {
	return []string{
		"qos",

		"policy",

		"shaper-hfsc",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfsc) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"shaper_hfsc_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Hierarchical Fair Service Curve's policy

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
    |  auto  &emsp; |  Bandwidth matches interface speed  |
    |  <number>  &emsp; |  Bits per second  |
    |  <number>bit  &emsp; |  Bits per second  |
    |  <number>kbit  &emsp; |  Kilobits per second  |
    |  <number>mbit  &emsp; |  Megabits per second  |
    |  <number>gbit  &emsp; |  Gigabits per second  |
    |  <number>tbit  &emsp; |  Terabits per second  |
    |  <number>%%  &emsp; |  Percentage of interface link speed  |

`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		// Nodes

		"default": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscDefault{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Default policy

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyShaperHfsc) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyShaperHfsc) UnmarshalJSON(_ []byte) error {
	return nil
}