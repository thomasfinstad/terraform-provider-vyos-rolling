// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnSstpAuthentication describes the resource data model.
type VpnSstpAuthentication struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafVpnSstpAuthenticationMode      types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafVpnSstpAuthenticationProtocols types.List   `tfsdk:"protocols" vyos:"protocols,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeVpnSstpAuthenticationLocalUsers bool `tfsdk:"-" vyos:"local-users,ignore,omitempty"`
	ExistsNodeVpnSstpAuthenticationRadius     bool `tfsdk:"-" vyos:"radius,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *VpnSstpAuthentication) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnSstpAuthentication) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"sstp",

		"authentication",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnSstpAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

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

		"protocols": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Authentication protocol for remote access peer SSTP VPN

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  pap  &emsp; |  Authentication via PAP (Password Authentication Protocol)  |
    |  chap  &emsp; |  Authentication via CHAP (Challenge Handshake Authentication Protocol)  |
    |  mschap  &emsp; |  Authentication via MS-CHAP (Microsoft Challenge Handshake Authentication Protocol)  |
    |  mschap-v2  &emsp; |  Authentication via MS-CHAPv2 (Microsoft Challenge Handshake Authentication Protocol, version 2)  |

`,

			// Default:          stringdefault.StaticString(`pap chap mschap mschap-v2`),
			Computed: true,
		},
	}
}