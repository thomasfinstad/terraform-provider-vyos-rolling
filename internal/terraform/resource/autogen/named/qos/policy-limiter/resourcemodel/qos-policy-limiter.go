// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyLimiter describes the resource data model.
type QosPolicyLimiter struct {
	SelfIdentifier types.String `tfsdk:"limiter_id" vyos:",self-id"`

	// LeafNodes
	LeafQosPolicyLimiterDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagQosPolicyLimiterClass bool `tfsdk:"class" vyos:"class,child"`

	// Nodes
	NodeQosPolicyLimiterDefault *QosPolicyLimiterDefault `tfsdk:"default" vyos:"default,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyLimiter) GetVyosPath() []string {
	return []string{
		"qos",

		"policy",

		"limiter",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyLimiter) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"limiter_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Traffic input limiting policy

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
			Attributes: QosPolicyLimiterDefault{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Default policy

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyLimiter) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyLimiter) UnmarshalJSON(_ []byte) error {
	return nil
}