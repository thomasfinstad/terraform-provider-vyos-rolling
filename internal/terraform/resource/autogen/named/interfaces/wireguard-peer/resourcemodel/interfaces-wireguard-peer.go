// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWireguardPeer describes the resource data model.
type InterfacesWireguardPeer struct {
	SelfIdentifier types.String `tfsdk:"peer_id" vyos:",self-id"`

	ParentIDInterfacesWireguard types.String `tfsdk:"wireguard" vyos:"wireguard,parent-id"`

	// LeafNodes
	LeafInterfacesWireguardPeerDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesWireguardPeerPublicKey           types.String `tfsdk:"public_key" vyos:"public-key,omitempty"`
	LeafInterfacesWireguardPeerPresharedKey        types.String `tfsdk:"preshared_key" vyos:"preshared-key,omitempty"`
	LeafInterfacesWireguardPeerAllowedIPs          types.List   `tfsdk:"allowed_ips" vyos:"allowed-ips,omitempty"`
	LeafInterfacesWireguardPeerAddress             types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesWireguardPeerPort                types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafInterfacesWireguardPeerPersistentKeepalive types.Number `tfsdk:"persistent_keepalive" vyos:"persistent-keepalive,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWireguardPeer) GetVyosPath() []string {
	return []string{
		"interfaces",

		"wireguard",
		o.ParentIDInterfacesWireguard.ValueString(),

		"peer",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWireguardPeer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"peer_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `peer alias

`,
		},

		"wireguard_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `WireGuard Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  wgN  &emsp; |  WireGuard interface name  |

`,
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"public_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `base64 encoded public key

`,
		},

		"preshared_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `base64 encoded preshared key

`,
		},

		"allowed_ips": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP addresses allowed to traverse the peer

`,
		},

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of tunnel endpoint

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address of remote tunnel endpoint  |
    |  ipv6  &emsp; |  IPv6 address of remote tunnel endpoint  |

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,
		},

		"persistent_keepalive": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval to send keepalive messages

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Interval in seconds  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWireguardPeer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWireguardPeer) UnmarshalJSON(_ []byte) error {
	return nil
}