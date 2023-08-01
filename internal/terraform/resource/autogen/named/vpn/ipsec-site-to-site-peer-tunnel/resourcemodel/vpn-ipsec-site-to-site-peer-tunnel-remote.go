// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnIPsecSiteToSitePeerTunnelRemote describes the resource data model.
type VpnIPsecSiteToSitePeerTunnelRemote struct {
	// LeafNodes
	LeafVpnIPsecSiteToSitePeerTunnelRemotePort   types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafVpnIPsecSiteToSitePeerTunnelRemotePrefix types.List   `tfsdk:"prefix" vyos:"prefix,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerTunnelRemote) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,
		},

		"prefix": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Remote IPv4 or IPv6 prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Remote IPv4 prefix  |
    |  ipv6net  &emsp; |  Remote IPv6 prefix  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecSiteToSitePeerTunnelRemote) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnIPsecSiteToSitePeerTunnelRemote) UnmarshalJSON(_ []byte) error {
	return nil
}
