// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessCapabilitiesHt describes the resource data model.
type InterfacesWirelessCapabilitiesHt struct {
	// LeafNodes
	LeafInterfacesWirelessCapabilitiesHtFourzeromhzIncapable types.Bool   `tfsdk:"40mhz_incapable" vyos:"40mhz-incapable,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtAutoPowersave        types.Bool   `tfsdk:"auto_powersave" vyos:"auto-powersave,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtChannelSetWIDth      types.List   `tfsdk:"channel_set_width" vyos:"channel-set-width,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtDelayedBlockAck      types.Bool   `tfsdk:"delayed_block_ack" vyos:"delayed-block-ack,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtDsssCckFourzero      types.Bool   `tfsdk:"dsss_cck_40" vyos:"dsss-cck-40,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtGreenfield           types.Bool   `tfsdk:"greenfield" vyos:"greenfield,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtLdpc                 types.Bool   `tfsdk:"ldpc" vyos:"ldpc,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtLsigProtection       types.Bool   `tfsdk:"lsig_protection" vyos:"lsig-protection,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtMaxAmsdu             types.String `tfsdk:"max_amsdu" vyos:"max-amsdu,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtShortGi              types.List   `tfsdk:"short_gi" vyos:"short-gi,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtSmps                 types.String `tfsdk:"smps" vyos:"smps,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesWirelessCapabilitiesHtStbc *InterfacesWirelessCapabilitiesHtStbc `tfsdk:"stbc" vyos:"stbc,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessCapabilitiesHt) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"40mhz_incapable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `40MHz intolerance, use 20MHz only!

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"auto_powersave": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable WMM-PS unscheduled automatic power aave delivery [U-APSD]

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"channel_set_width": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Supported channel set width

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ht20  &emsp; |  Supported channel set width both 20 MHz only  |
    |  ht40+  &emsp; |  Supported channel set width both 20 MHz and 40 MHz with secondary channel above primary channel  |
    |  ht40-  &emsp; |  Supported channel set width both 20 MHz and 40 MHz with secondary channel below primary channel  |

`,
		},

		"delayed_block_ack": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable HT-delayed block ack

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"dsss_cck_40": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable DSSS_CCK-40

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"greenfield": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable HT-greenfield

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ldpc": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable LDPC coding capability

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"lsig_protection": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable L-SIG TXOP protection capability

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"max_amsdu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set maximum A-MSDU length

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  3839  &emsp; |  Set maximum A-MSDU length to 3839 octets  |
    |  7935  &emsp; |  Set maximum A-MSDU length to 7935 octets  |

`,
		},

		"short_gi": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Short GI capabilities

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  20  &emsp; |  Short GI for 20 MHz  |
    |  40  &emsp; |  Short GI for 40 MHz  |

`,
		},

		"smps": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Spatial Multiplexing Power Save (SMPS) settings

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  static  &emsp; |  STATIC Spatial Multiplexing (SM) Power Save  |
    |  dynamic  &emsp; |  DYNAMIC Spatial Multiplexing (SM) Power Save  |

`,
		},

		// Nodes

		"stbc": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesHtStbc{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Support for sending and receiving PPDU using STBC (Space Time Block Coding)

`,
		},
	}
}
