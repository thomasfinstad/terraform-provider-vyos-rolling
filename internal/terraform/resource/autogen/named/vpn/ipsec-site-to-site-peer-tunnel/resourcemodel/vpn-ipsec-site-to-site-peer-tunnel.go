// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnIPsecSiteToSitePeerTunnel describes the resource data model.
type VpnIPsecSiteToSitePeerTunnel struct {
	SelfIdentifier types.Number `tfsdk:"tunnel_id" vyos:",self-id"`

	ParentIDVpnIPsecSiteToSitePeer types.String `tfsdk:"peer" vyos:"peer,parent-id"`

	// LeafNodes
	LeafVpnIPsecSiteToSitePeerTunnelDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnIPsecSiteToSitePeerTunnelEspGroup types.String `tfsdk:"esp_group" vyos:"esp-group,omitempty"`
	LeafVpnIPsecSiteToSitePeerTunnelProtocol types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafVpnIPsecSiteToSitePeerTunnelPriority types.Number `tfsdk:"priority" vyos:"priority,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVpnIPsecSiteToSitePeerTunnelLocal  *VpnIPsecSiteToSitePeerTunnelLocal  `tfsdk:"local" vyos:"local,omitempty"`
	NodeVpnIPsecSiteToSitePeerTunnelRemote *VpnIPsecSiteToSitePeerTunnelRemote `tfsdk:"remote" vyos:"remote,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecSiteToSitePeerTunnel) GetVyosPath() []string {
	return []string{
		"vpn",

		"ipsec",

		"site-to-site",

		"peer",
		o.ParentIDVpnIPsecSiteToSitePeer.ValueString(),

		"tunnel",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerTunnel) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `tunnel_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"tunnel_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Peer tunnel

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Peer tunnel  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"peer_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Connection name of the peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Connection name of the peer  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"esp_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encapsulating Security Payloads (ESP) group name

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Protocol name  |

`,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Priority for IPsec policy (lowest value more preferable)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-100  &emsp; |  Priority for IPsec policy (lowest value more preferable)  |

`,
		},

		// Nodes

		"local": schema.SingleNestedAttribute{
			Attributes: VpnIPsecSiteToSitePeerTunnelLocal{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Local parameters for interesting traffic

`,
		},

		"remote": schema.SingleNestedAttribute{
			Attributes: VpnIPsecSiteToSitePeerTunnelRemote{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match remote addresses

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecSiteToSitePeerTunnel) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnIPsecSiteToSitePeerTunnel) UnmarshalJSON(_ []byte) error {
	return nil
}
