// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfAreaRange describes the resource data model.
type VrfNameProtocolsOspfAreaRange struct {
	// LeafNodes
	VrfNameProtocolsOspfAreaRangeCost         customtypes.CustomStringValue `tfsdk:"cost" json:"cost,omitempty"`
	VrfNameProtocolsOspfAreaRangeNotAdvertise customtypes.CustomStringValue `tfsdk:"not_advertise" json:"not-advertise,omitempty"`
	VrfNameProtocolsOspfAreaRangeSubstitute   customtypes.CustomStringValue `tfsdk:"substitute" json:"substitute,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaRange) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cost": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Metric for this range

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777215  |  Metric for this range  |

`,
		},

		"not_advertise": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Do not advertise this range

`,
		},

		"substitute": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Advertise area range as another prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Advertise area range as another prefix  |

`,
		},

		// TagNodes

		// Nodes

	}
}
