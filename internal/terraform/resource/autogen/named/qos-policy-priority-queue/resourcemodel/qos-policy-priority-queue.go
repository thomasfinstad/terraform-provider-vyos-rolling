// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyPriorityQueue describes the resource data model.
type QosPolicyPriorityQueue struct {
	// LeafNodes
	QosPolicyPriorityQueueDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`

	// TagNodes
	QosPolicyPriorityQueueClass types.Map `tfsdk:"class" json:"class,omitempty"`

	// Nodes
	QosPolicyPriorityQueueDefault types.Object `tfsdk:"default" json:"default,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyPriorityQueue) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		// TagNodes

		"class": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: QosPolicyPriorityQueueClass{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Class Handle

|  Format  |  Description  |
|----------|---------------|
|  u32:1-7  |  Priority  |

`,
		},

		// Nodes

		"default": schema.SingleNestedAttribute{
			Attributes: QosPolicyPriorityQueueDefault{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Default policy

`,
		},
	}
}
