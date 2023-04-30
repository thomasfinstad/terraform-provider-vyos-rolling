// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyRandomDetect describes the resource data model.
type QosPolicyRandomDetect struct {
	// LeafNodes
	QosPolicyRandomDetectDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	QosPolicyRandomDetectBandwIDth   customtypes.CustomStringValue `tfsdk:"bandwidth" json:"bandwidth,omitempty"`

	// TagNodes
	QosPolicyRandomDetectPrecedence types.Map `tfsdk:"precedence" json:"precedence,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyRandomDetect) ResourceAttributes() map[string]schema.Attribute {
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

		"bandwidth": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Available bandwidth for this policy

|  Format  |  Description  |
|----------|---------------|
|  auto  |  Bandwidth matches interface speed  |
|  <number>  |  Bits per second  |
|  <number>bit  |  Bits per second  |
|  <number>kbit  |  Kilobits per second  |
|  <number>mbit  |  Megabits per second  |
|  <number>gbit  |  Gigabits per second  |
|  <number>tbit  |  Terabits per second  |
|  <number>%%  |  Percentage of interface link speed  |

`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		// TagNodes

		"precedence": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: QosPolicyRandomDetectPrecedence{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `IP precedence

|  Format  |  Description  |
|----------|---------------|
|  u32:0-7  |  IP precedence value  |

`,
		},

		// Nodes

	}
}
