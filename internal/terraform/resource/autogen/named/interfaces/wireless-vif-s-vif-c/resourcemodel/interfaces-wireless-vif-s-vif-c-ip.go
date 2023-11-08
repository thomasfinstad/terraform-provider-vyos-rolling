// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessVifSVifCIP describes the resource data model.
type InterfacesWirelessVifSVifCIP struct {
	// LeafNodes
	LeafInterfacesWirelessVifSVifCIPAdjustMss               types.String `tfsdk:"adjust_mss" vyos:"adjust-mss,omitempty"`
	LeafInterfacesWirelessVifSVifCIPArpCacheTimeout         types.Number `tfsdk:"arp_cache_timeout" vyos:"arp-cache-timeout,omitempty"`
	LeafInterfacesWirelessVifSVifCIPDisableArpFilter        types.Bool   `tfsdk:"disable_arp_filter" vyos:"disable-arp-filter,omitempty"`
	LeafInterfacesWirelessVifSVifCIPDisableForwarding       types.Bool   `tfsdk:"disable_forwarding" vyos:"disable-forwarding,omitempty"`
	LeafInterfacesWirelessVifSVifCIPEnableDirectedBroadcast types.Bool   `tfsdk:"enable_directed_broadcast" vyos:"enable-directed-broadcast,omitempty"`
	LeafInterfacesWirelessVifSVifCIPEnableArpAccept         types.Bool   `tfsdk:"enable_arp_accept" vyos:"enable-arp-accept,omitempty"`
	LeafInterfacesWirelessVifSVifCIPEnableArpAnnounce       types.Bool   `tfsdk:"enable_arp_announce" vyos:"enable-arp-announce,omitempty"`
	LeafInterfacesWirelessVifSVifCIPEnableArpIgnore         types.Bool   `tfsdk:"enable_arp_ignore" vyos:"enable-arp-ignore,omitempty"`
	LeafInterfacesWirelessVifSVifCIPEnableProxyArp          types.Bool   `tfsdk:"enable_proxy_arp" vyos:"enable-proxy-arp,omitempty"`
	LeafInterfacesWirelessVifSVifCIPProxyArpPvlan           types.Bool   `tfsdk:"proxy_arp_pvlan" vyos:"proxy-arp-pvlan,omitempty"`
	LeafInterfacesWirelessVifSVifCIPSourceValIDation        types.String `tfsdk:"source_validation" vyos:"source-validation,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifSVifCIP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adjust TCP MSS value

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  clamp-mss-to-pmtu  &emsp; |  Automatically sets the MSS to the proper value  |
    |  number: 536-65535  &emsp; |  TCP Maximum segment size in bytes  |

`,
		},

		"arp_cache_timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ARP cache entry timeout in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-86400  &emsp; |  ARP cache entry timout in seconds  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"disable_arp_filter": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable ARP filter on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_forwarding": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_directed_broadcast": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable directed broadcast forwarding on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_arp_accept": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable ARP accept on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_arp_announce": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable ARP announce on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_arp_ignore": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable ARP ignore on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_proxy_arp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable proxy-arp on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"proxy_arp_pvlan": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable private VLAN proxy ARP on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"source_validation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source validation by reversed path (RFC3704)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  strict  &emsp; |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose  &emsp; |  Enable Loose Reverse Path Forwarding as defined in RFC3704  |
    |  disable  &emsp; |  No source validation  |

`,
		},

		// Nodes

	}
}
