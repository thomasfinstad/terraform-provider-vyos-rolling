// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance struct {
	// LeafNodes
	VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceExternal customtypes.CustomStringValue `tfsdk:"external" json:"external,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceInternal customtypes.CustomStringValue `tfsdk:"internal" json:"internal,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistanceLocal    customtypes.CustomStringValue `tfsdk:"local" json:"local,omitempty"`

	// TagNodes
	VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix types.Map `tfsdk:"prefix" json:"prefix,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `eBGP routes administrative distance

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  eBGP routes administrative distance  |

`,
		},

		"internal": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `iBGP routes administrative distance

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  iBGP routes administrative distance  |

`,
		},

		"local": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Locally originated BGP routes administrative distance

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Locally originated BGP routes administrative distance  |

`,
		},

		// TagNodes

		"prefix": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Administrative distance for a specific BGP prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Administrative distance for a specific BGP prefix  |

`,
		},

		// Nodes

	}
}
