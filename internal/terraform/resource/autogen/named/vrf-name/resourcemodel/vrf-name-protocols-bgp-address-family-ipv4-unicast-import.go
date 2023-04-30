// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastImport describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastImport struct {
	// LeafNodes
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastImportVpn customtypes.CustomStringValue `tfsdk:"vpn" json:"vpn,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastImportVrf customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastImport) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"vpn": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `to/from default instance VPN RIB

`,
		},

		"vrf": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VRF to import from

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

`,
		},

		// TagNodes

		// Nodes

	}
}
