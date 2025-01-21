/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (wireless) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &InterfacesWireless{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (wireless) */
// InterfacesWirelessSelfIdentifier used by TagNodes to keep the id info
type InterfacesWirelessSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (wireless) */

	InterfacesWireless types.String `tfsdk:"wireless" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interfaces) */
}

// InterfacesWireless describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type InterfacesWireless struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *InterfacesWirelessSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafInterfacesWirelessAddress                types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesWirelessChannel                types.String `tfsdk:"channel" vyos:"channel,omitempty"`
	LeafInterfacesWirelessDescrIPtion            types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesWirelessDisableBroadcastSsID   types.Bool   `tfsdk:"disable_broadcast_ssid" vyos:"disable-broadcast-ssid,omitempty"`
	LeafInterfacesWirelessDisableLinkDetect      types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesWirelessDisable                types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesWirelessVrf                    types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafInterfacesWirelessExpungeFailingStations types.Bool   `tfsdk:"expunge_failing_stations" vyos:"expunge-failing-stations,omitempty"`
	LeafInterfacesWirelessHwID                   types.String `tfsdk:"hw_id" vyos:"hw-id,omitempty"`
	LeafInterfacesWirelessIsolateStations        types.Bool   `tfsdk:"isolate_stations" vyos:"isolate-stations,omitempty"`
	LeafInterfacesWirelessMac                    types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesWirelessMaxStations            types.Number `tfsdk:"max_stations" vyos:"max-stations,omitempty"`
	LeafInterfacesWirelessStationaryAp           types.Bool   `tfsdk:"stationary_ap" vyos:"stationary-ap,omitempty"`
	LeafInterfacesWirelessMgmtFrameProtection    types.String `tfsdk:"mgmt_frame_protection" vyos:"mgmt-frame-protection,omitempty"`
	LeafInterfacesWirelessEnableBfProtection     types.Bool   `tfsdk:"enable_bf_protection" vyos:"enable-bf-protection,omitempty"`
	LeafInterfacesWirelessMode                   types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafInterfacesWirelessPhysicalDevice         types.String `tfsdk:"physical_device" vyos:"physical-device,omitempty"`
	LeafInterfacesWirelessReduceTransmitPower    types.Number `tfsdk:"reduce_transmit_power" vyos:"reduce-transmit-power,omitempty"`
	LeafInterfacesWirelessSsID                   types.String `tfsdk:"ssid" vyos:"ssid,omitempty"`
	LeafInterfacesWirelessBssID                  types.String `tfsdk:"bssid" vyos:"bssid,omitempty"`
	LeafInterfacesWirelessType                   types.String `tfsdk:"type" vyos:"type,omitempty"`
	LeafInterfacesWirelessPerClientThread        types.Bool   `tfsdk:"per_client_thread" vyos:"per-client-thread,omitempty"`
	LeafInterfacesWirelessRedirect               types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`

	// TagNodes

	ExistsTagInterfacesWirelessVif bool `tfsdk:"-" vyos:"vif,child"`

	ExistsTagInterfacesWirelessVifS bool `tfsdk:"-" vyos:"vif-s,child"`

	// Nodes

	NodeInterfacesWirelessCapabilities *InterfacesWirelessCapabilities `tfsdk:"capabilities" vyos:"capabilities,omitempty"`

	NodeInterfacesWirelessDhcpOptions *InterfacesWirelessDhcpOptions `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`

	NodeInterfacesWirelessDhcpvsixOptions *InterfacesWirelessDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`

	NodeInterfacesWirelessIP *InterfacesWirelessIP `tfsdk:"ip" vyos:"ip,omitempty"`

	NodeInterfacesWirelessIPvsix *InterfacesWirelessIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`

	NodeInterfacesWirelessMirror *InterfacesWirelessMirror `tfsdk:"mirror" vyos:"mirror,omitempty"`

	NodeInterfacesWirelessSecURIty *InterfacesWirelessSecURIty `tfsdk:"security" vyos:"security,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesWireless) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesWireless) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesWireless) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWireless) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"wireless",
		o.SelfIdentifier.InterfacesWireless.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesWireless) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (interfaces) */
		"interfaces", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesWireless) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (interfaces) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWireless) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"wireless": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Wireless (WiFi/WLAN) Network Interface

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  wlanN   |  Wireless (WiFi/WLAN) interface name  |
`,
					Description: `Wireless (WiFi/WLAN) Network Interface

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  wlanN   |  Wireless (WiFi/WLAN) interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in wireless, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  wireless, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (interfaces) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address

    |  Format   |  Description                                   |
    |-----------|------------------------------------------------|
    |  ipv4net  |  IPv4 address and prefix length                |
    |  ipv6net  |  IPv6 address and prefix length                |
    |  dhcp     |  Dynamic Host Configuration Protocol           |
    |  dhcpv6   |  Dynamic Host Configuration Protocol for IPv6  |
`,
			Description: `IP address

    |  Format   |  Description                                   |
    |-----------|------------------------------------------------|
    |  ipv4net  |  IPv4 address and prefix length                |
    |  ipv6net  |  IPv6 address and prefix length                |
    |  dhcp     |  Dynamic Host Configuration Protocol           |
    |  dhcpv6   |  Dynamic Host Configuration Protocol for IPv6  |
`,
		},

		"channel":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (channel) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Wireless radio channel

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  0       |  Automatic Channel Selection (ACS)  |
    |  1-14    |  2.4Ghz (802.11 b/g/n/ax) Channel   |
    |  34-177  |  5Ghz (802.11 a/h/j/n/ac) Channel   |
    |  1-233   |  6Ghz (802.11 ax) Channel           |
`,
			Description: `Wireless radio channel

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  0       |  Automatic Channel Selection (ACS)  |
    |  1-14    |  2.4Ghz (802.11 b/g/n/ax) Channel   |
    |  34-177  |  5Ghz (802.11 a/h/j/n/ac) Channel   |
    |  1-233   |  6Ghz (802.11 ax) Channel           |
`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (description) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"disable_broadcast_ssid":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable-broadcast-ssid) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable broadcast of SSID from access-point

`,
			Description: `Disable broadcast of SSID from access-point

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_link_detect":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable-link-detect) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore link state changes

`,
			Description: `Ignore link state changes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Description: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vrf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (vrf) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
		},

		"expunge_failing_stations":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (expunge-failing-stations) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disassociate stations based on excessive transmission failures

`,
			Description: `Disassociate stations based on excessive transmission failures

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hw_id":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hw-id) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Associate Ethernet Interface with given Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
			Description: `Associate Ethernet Interface with given Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
		},

		"isolate_stations":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (isolate-stations) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Isolate stations on the AP so they cannot see each other

`,
			Description: `Isolate stations on the AP so they cannot see each other

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mac":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mac) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
			Description: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
		},

		"max_stations":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (max-stations) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of wireless radio stations. Excess stations will be rejected upon authentication request.

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-2007  |  Number of allowed stations  |
`,
			Description: `Maximum number of wireless radio stations. Excess stations will be rejected upon authentication request.

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-2007  |  Number of allowed stations  |
`,
		},

		"stationary_ap":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (stationary-ap) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Stationary AP config indicates that the AP doesn't move.

`,
			Description: `Stationary AP config indicates that the AP doesn't move.

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mgmt_frame_protection":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mgmt-frame-protection) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Management Frame Protection (MFP) according to IEEE 802.11w

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  disabled  |  no MFP                             |
    |  optional  |  MFP optional                       |
    |  required  |  MFP enforced (mandatory for WPA3)  |
`,
			Description: `Management Frame Protection (MFP) according to IEEE 802.11w

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  disabled  |  no MFP                             |
    |  optional  |  MFP optional                       |
    |  required  |  MFP enforced (mandatory for WPA3)  |
`,

			// Default:          stringdefault.StaticString(`disabled`),
			Computed: true,
		},

		"enable_bf_protection":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (enable-bf-protection) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Beacon Protection: management frame protection for Beacon frames, requires Management Frame Protection (MFP)

`,
			Description: `Beacon Protection: management frame protection for Beacon frames, requires Management Frame Protection (MFP)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mode":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mode) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Wireless radio mode

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  a       |  802.11a - 54 Mbits/sec        |
    |  b       |  802.11b - 11 Mbits/sec        |
    |  g       |  802.11g - 54 Mbits/sec        |
    |  n       |  802.11n - 600 Mbits/sec       |
    |  ac      |  802.11ac - 1300 Mbits/sec     |
    |  ax      |  802.11ax (6GHz only for now)  |
`,
			Description: `Wireless radio mode

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  a       |  802.11a - 54 Mbits/sec        |
    |  b       |  802.11b - 11 Mbits/sec        |
    |  g       |  802.11g - 54 Mbits/sec        |
    |  n       |  802.11n - 600 Mbits/sec       |
    |  ac      |  802.11ac - 1300 Mbits/sec     |
    |  ax      |  802.11ax (6GHz only for now)  |
`,

			// Default:          stringdefault.StaticString(`g`),
			Computed: true,
		},

		"physical_device":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (physical-device) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Wireless physical device

`,
			Description: `Wireless physical device

`,

			// Default:          stringdefault.StaticString(`phy0`),
			Computed: true,
		},

		"reduce_transmit_power":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (reduce-transmit-power) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Transmission power reduction in dBm

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  0-255   |  TX power reduction in dBm  |
`,
			Description: `Transmission power reduction in dBm

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  0-255   |  TX power reduction in dBm  |
`,
		},

		"ssid":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ssid) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Wireless access-point service set identifier (SSID)

`,
			Description: `Wireless access-point service set identifier (SSID)

`,
		},

		"bssid":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (bssid) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Basic Service Set Identifier (BSSID) - currently station mode only

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  macaddr  |  BSSID (MAC) address  |
`,
			Description: `Basic Service Set Identifier (BSSID) - currently station mode only

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  macaddr  |  BSSID (MAC) address  |
`,
		},

		"type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (type) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Wireless device type for this interface

    |  Format        |  Description                                             |
    |----------------|----------------------------------------------------------|
    |  access-point  |  Access-point forwards packets between other nodes       |
    |  station       |  Connects to another access point                        |
    |  monitor       |  Passively monitor all packets on the frequency/channel  |
`,
			Description: `Wireless device type for this interface

    |  Format        |  Description                                             |
    |----------------|----------------------------------------------------------|
    |  access-point  |  Access-point forwards packets between other nodes       |
    |  station       |  Connects to another access point                        |
    |  monitor       |  Passively monitor all packets on the frequency/channel  |
`,

			// Default:          stringdefault.StaticString(`monitor`),
			Computed: true,
		},

		"per_client_thread":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (per-client-thread) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Process traffic from each client in a dedicated thread

`,
			Description: `Process traffic from each client in a dedicated thread

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"redirect":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (redirect) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
			Description: `Redirect incoming packet to destination

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
		},

		// TagNodes

		// Nodes

		"capabilities": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilities{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `HT and VHT capabilities for your card

`,
			Description: `HT and VHT capabilities for your card

`,
		},

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessDhcpOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
			Description: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessDhcpvsixOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
			Description: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessMirror{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
			Description: `Mirror ingress/egress packets

`,
		},

		"security": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURIty{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Wireless security settings

`,
			Description: `Wireless security settings

`,
		},
	}
}
