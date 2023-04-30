// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork struct {
	// LeafNodes
	VrfNameProtocolsBgpAddressFamilyIPvsixVpnNetworkRd    customtypes.CustomStringValue `tfsdk:"rd" json:"rd,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvsixVpnNetworkLabel customtypes.CustomStringValue `tfsdk:"label" json:"label,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"rd": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route Distinguisher

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |

`,
		},

		"label": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `MPLS label value assigned to route

|  Format  |  Description  |
|----------|---------------|
|  u32:0-1048575  |  MPLS label value  |

`,
		},

		// TagNodes

		// Nodes

	}
}
