// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesWireless describes the resource data model.
type InterfacesWireless struct {
	// LeafNodes
	InterfacesWirelessAddress                customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	InterfacesWirelessChannel                customtypes.CustomStringValue `tfsdk:"channel" json:"channel,omitempty"`
	InterfacesWirelessCountryCode            customtypes.CustomStringValue `tfsdk:"country_code" json:"country-code,omitempty"`
	InterfacesWirelessDescrIPtion            customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	InterfacesWirelessDisableBroadcastSsID   customtypes.CustomStringValue `tfsdk:"disable_broadcast_ssid" json:"disable-broadcast-ssid,omitempty"`
	InterfacesWirelessDisableLinkDetect      customtypes.CustomStringValue `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	InterfacesWirelessDisable                customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	InterfacesWirelessVrf                    customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`
	InterfacesWirelessExpungeFailingStations customtypes.CustomStringValue `tfsdk:"expunge_failing_stations" json:"expunge-failing-stations,omitempty"`
	InterfacesWirelessHwID                   customtypes.CustomStringValue `tfsdk:"hw_id" json:"hw-id,omitempty"`
	InterfacesWirelessIsolateStations        customtypes.CustomStringValue `tfsdk:"isolate_stations" json:"isolate-stations,omitempty"`
	InterfacesWirelessMac                    customtypes.CustomStringValue `tfsdk:"mac" json:"mac,omitempty"`
	InterfacesWirelessMaxStations            customtypes.CustomStringValue `tfsdk:"max_stations" json:"max-stations,omitempty"`
	InterfacesWirelessMgmtFrameProtection    customtypes.CustomStringValue `tfsdk:"mgmt_frame_protection" json:"mgmt-frame-protection,omitempty"`
	InterfacesWirelessMode                   customtypes.CustomStringValue `tfsdk:"mode" json:"mode,omitempty"`
	InterfacesWirelessPhysicalDevice         customtypes.CustomStringValue `tfsdk:"physical_device" json:"physical-device,omitempty"`
	InterfacesWirelessReduceTransmitPower    customtypes.CustomStringValue `tfsdk:"reduce_transmit_power" json:"reduce-transmit-power,omitempty"`
	InterfacesWirelessSsID                   customtypes.CustomStringValue `tfsdk:"ssid" json:"ssid,omitempty"`
	InterfacesWirelessType                   customtypes.CustomStringValue `tfsdk:"type" json:"type,omitempty"`
	InterfacesWirelessRedirect               customtypes.CustomStringValue `tfsdk:"redirect" json:"redirect,omitempty"`

	// TagNodes
	InterfacesWirelessVif  types.Map `tfsdk:"vif" json:"vif,omitempty"`
	InterfacesWirelessVifS types.Map `tfsdk:"vif_s" json:"vif-s,omitempty"`

	// Nodes
	InterfacesWirelessCapabilities    types.Object `tfsdk:"capabilities" json:"capabilities,omitempty"`
	InterfacesWirelessDhcpOptions     types.Object `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	InterfacesWirelessDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	InterfacesWirelessIP              types.Object `tfsdk:"ip" json:"ip,omitempty"`
	InterfacesWirelessIPvsix          types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
	InterfacesWirelessMirror          types.Object `tfsdk:"mirror" json:"mirror,omitempty"`
	InterfacesWirelessSecURIty        types.Object `tfsdk:"security" json:"security,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesWireless) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |
|  dhcp  |  Dynamic Host Configuration Protocol  |
|  dhcpv6  |  Dynamic Host Configuration Protocol for IPv6  |

`,
		},

		"channel": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Wireless radio channel

|  Format  |  Description  |
|----------|---------------|
|  0  |  Automatic Channel Selection (ACS)  |
|  u32:1-14  |  2.4Ghz (802.11 b/g/n) Channel  |
|  u32:34-173  |  5Ghz (802.11 a/h/j/n/ac) Channel  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"country_code": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Indicate country in which device is operating

|  Format  |  Description  |
|----------|---------------|
|  txt  |  ISO/IEC 3166-1 Country Code  |

`,
		},

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"disable_broadcast_ssid": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable broadcast of SSID from access-point

`,
		},

		"disable_link_detect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Ignore link state changes

`,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"vrf": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

`,
		},

		"expunge_failing_stations": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disassociate stations based on excessive transmission failures

`,
		},

		"hw_id": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Associate Ethernet Interface with given Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |

`,
		},

		"isolate_stations": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Isolate stations on the AP so they cannot see each other

`,
		},

		"mac": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |

`,
		},

		"max_stations": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum number of wireless radio stations. Excess stations will be rejected upon authentication request.

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2007  |  Number of allowed stations  |

`,
		},

		"mgmt_frame_protection": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Management Frame Protection (MFP) according to IEEE 802.11w

|  Format  |  Description  |
|----------|---------------|
|  disabled  |  no MFP  |
|  optional  |  MFP optional  |
|  required  |  MFP enforced  |

`,

			// Default:          stringdefault.StaticString(`disabled`),
			Computed: true,
		},

		"mode": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Wireless radio mode

|  Format  |  Description  |
|----------|---------------|
|  a  |  802.11a - 54 Mbits/sec  |
|  b  |  802.11b - 11 Mbits/sec  |
|  g  |  802.11g - 54 Mbits/sec  |
|  n  |  802.11n - 600 Mbits/sec  |
|  ac  |  802.11ac - 1300 Mbits/sec  |

`,

			// Default:          stringdefault.StaticString(`g`),
			Computed: true,
		},

		"physical_device": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Wireless physical device

`,

			// Default:          stringdefault.StaticString(`phy0`),
			Computed: true,
		},

		"reduce_transmit_power": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Transmission power reduction in dBm

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  TX power reduction in dBm  |

`,
		},

		"ssid": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Wireless access-point service set identifier (SSID)

`,
		},

		"type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Wireless device type for this interface

|  Format  |  Description  |
|----------|---------------|
|  access-point  |  Access-point forwards packets between other nodes  |
|  station  |  Connects to another access point  |
|  monitor  |  Passively monitor all packets on the frequency/channel  |

`,

			// Default:          stringdefault.StaticString(`monitor`),
			Computed: true,
		},

		"redirect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |

`,
		},

		// TagNodes

		"vif": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesWirelessVif{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4094  |  Virtual Local Area Network (VLAN) ID  |

`,
		},

		"vif_s": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesWirelessVifS{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4094  |  QinQ Virtual Local Area Network (VLAN) ID  |

`,
		},

		// Nodes

		"capabilities": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilities{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `HT and VHT capabilities for your card

`,
		},

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessDhcpOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessDhcpvsixOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessMirror{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"security": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURIty{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Wireless security settings

`,
		},
	}
}
