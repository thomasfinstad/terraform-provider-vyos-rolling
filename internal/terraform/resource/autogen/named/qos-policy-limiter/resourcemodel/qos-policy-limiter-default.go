// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyLimiterDefault describes the resource data model.
type QosPolicyLimiterDefault struct {
	// LeafNodes
	QosPolicyLimiterDefaultBandwIDth customtypes.CustomStringValue `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
	QosPolicyLimiterDefaultBurst     customtypes.CustomStringValue `tfsdk:"burst" json:"burst,omitempty"`
	QosPolicyLimiterDefaultExceed    customtypes.CustomStringValue `tfsdk:"exceed" json:"exceed,omitempty"`
	QosPolicyLimiterDefaultNotExceed customtypes.CustomStringValue `tfsdk:"not_exceed" json:"not-exceed,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyLimiterDefault) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bandwidth": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Available bandwidth for this policy

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bits per second  |
|  <number>bit  |  Bits per second  |
|  <number>kbit  |  Kilobits per second  |
|  <number>mbit  |  Megabits per second  |
|  <number>gbit  |  Gigabits per second  |
|  <number>tbit  |  Terabits per second  |
|  <number>%  |  Percentage of interface link speed  |

`,
		},

		"burst": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Burst size for this class

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bytes  |
|  <number><suffix>  |  Bytes with scaling suffix (kb, mb, gb)  |

`,

			// Default:          stringdefault.StaticString(`15k`),
			Computed: true,
		},

		"exceed": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Default action for packets exceeding the limiter

|  Format  |  Description  |
|----------|---------------|
|  continue  |  Do not do anything, just continue with the next action in line  |
|  drop  |  Drop the packet immediately  |
|  ok  |  Accept the packet  |
|  reclassify  |  Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
|  pipe  |  Pass the packet to the next action in line  |

`,

			// Default:          stringdefault.StaticString(`drop`),
			Computed: true,
		},

		"not_exceed": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Default action for packets not exceeding the limiter

|  Format  |  Description  |
|----------|---------------|
|  continue  |  Do not do anything, just continue with the next action in line  |
|  drop  |  Drop the packet immediately  |
|  ok  |  Accept the packet  |
|  reclassify  |  Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
|  pipe  |  Pass the packet to the next action in line  |

`,

			// Default:          stringdefault.StaticString(`ok`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
