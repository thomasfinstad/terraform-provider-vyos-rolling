// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnPptpRemoteAccessAuthentication describes the resource data model.
type VpnPptpRemoteAccessAuthentication struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafVpnPptpRemoteAccessAuthenticationRequire types.String `tfsdk:"require" vyos:"require,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationMppe    types.String `tfsdk:"mppe" vyos:"mppe,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationMode    types.String `tfsdk:"mode" vyos:"mode,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeVpnPptpRemoteAccessAuthenticationLocalUsers bool `tfsdk:"-" vyos:"local-users,ignore,omitempty"`
	ExistsNodeVpnPptpRemoteAccessAuthenticationRadius     bool `tfsdk:"-" vyos:"radius,ignore,omitempty"`
	ExistsNodeVpnPptpRemoteAccessAuthenticationRateLimit  bool `tfsdk:"-" vyos:"rate-limit,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *VpnPptpRemoteAccessAuthentication) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnPptpRemoteAccessAuthentication) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"pptp",

		"remote-access",

		"authentication",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnPptpRemoteAccessAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"require": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication protocol for remote access peer PPTP VPN

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  pap  &emsp; |  Require the peer to authenticate itself using PAP [Password Authentication Protocol].  |
    |  chap  &emsp; |  Require the peer to authenticate itself using CHAP [Challenge Handshake Authentication Protocol].  |
    |  mschap  &emsp; |  Require the peer to authenticate itself using CHAP [Challenge Handshake Authentication Protocol].  |
    |  mschap-v2  &emsp; |  Require the peer to authenticate itself using MS-CHAPv2 [Microsoft Challenge Handshake Authentication Protocol, Version 2].  |

`,
		},

		"mppe": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies mppe negotioation preference. (default require mppe 128-bit stateless

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  deny  &emsp; |  deny mppe  |
    |  prefer  &emsp; |  ask client for mppe, if it rejects do not fail  |
    |  require  &emsp; |  ask client for mppe, if it rejects drop connection  |

`,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication mode used by this server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  local  &emsp; |  Use local username/password configuration  |
    |  radius  &emsp; |  Use RADIUS server for user autentication  |
    |  noauth  &emsp; |  Authentication disabled  |

`,

			// Default:          stringdefault.StaticString(`local`),
			Computed: true,
		},
	}
}