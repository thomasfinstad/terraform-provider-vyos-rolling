// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyShaperHfsc describes the resource data model.
type QosPolicyShaperHfsc struct {
	// LeafNodes
	QosPolicyShaperHfscDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	QosPolicyShaperHfscBandwIDth   customtypes.CustomStringValue `tfsdk:"bandwidth" json:"bandwidth,omitempty"`

	// TagNodes
	QosPolicyShaperHfscClass types.Map `tfsdk:"class" json:"class,omitempty"`

	// Nodes
	QosPolicyShaperHfscDefault types.Object `tfsdk:"default" json:"default,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyShaperHfsc) ResourceAttributes() map[string]schema.Attribute {
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

		"class": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: QosPolicyShaperHfscClass{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Class ID

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4095  |  Class Identifier  |

`,
		},

		// Nodes

		"default": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscDefault{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Default policy

`,
		},
	}
}
