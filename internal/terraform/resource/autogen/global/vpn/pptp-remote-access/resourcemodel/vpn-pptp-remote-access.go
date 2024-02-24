// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnPptpRemoteAccess describes the resource data model.
type VpnPptpRemoteAccess struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafVpnPptpRemoteAccessMtu            types.String `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafVpnPptpRemoteAccessOutsIDeAddress types.String `tfsdk:"outside_address" vyos:"outside-address,omitempty"`
	LeafVpnPptpRemoteAccessNameServer     types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`
	LeafVpnPptpRemoteAccessWinsServer     types.List   `tfsdk:"wins_server" vyos:"wins-server,omitempty"`
	LeafVpnPptpRemoteAccessGatewayAddress types.String `tfsdk:"gateway_address" vyos:"gateway-address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeVpnPptpRemoteAccessClientIPPool   bool `tfsdk:"-" vyos:"client-ip-pool,ignore,omitempty"`
	ExistsNodeVpnPptpRemoteAccessAuthentication bool `tfsdk:"-" vyos:"authentication,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *VpnPptpRemoteAccess) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnPptpRemoteAccess) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"pptp",

		"remote-access",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnPptpRemoteAccess) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU) - default 1492

`,

			// Default:          stringdefault.StaticString(`1492`),
			Computed: true,
		},

		"outside_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `External IP address to which VPN clients will connect

`,
		},

		"name_server": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Domain Name Server (DNS) IPv4 address  |

`,
		},

		"wins_server": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Windows Internet Name Service (WINS) servers propagated to client

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Domain Name Server (DNS) IPv4 address  |

`,
		},

		"gateway_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Default Gateway send to the client  |

`,
		},
	}
}
