// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VpnIPsecProfileBind describes the resource data model.
type VpnIPsecProfileBind struct {
	// LeafNodes
	VpnIPsecProfileBindTunnel customtypes.CustomStringValue `tfsdk:"tunnel" json:"tunnel,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VpnIPsecProfileBind) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"tunnel": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Tunnel interface associated with this profile

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Associated interface to this profile  |

`,
		},

		// TagNodes

		// Nodes

	}
}
