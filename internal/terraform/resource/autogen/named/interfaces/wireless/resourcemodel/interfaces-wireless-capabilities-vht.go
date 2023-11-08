// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessCapabilitiesVht describes the resource data model.
type InterfacesWirelessCapabilitiesVht struct {
	// LeafNodes
	LeafInterfacesWirelessCapabilitiesVhtAntennaCount        types.Number `tfsdk:"antenna_count" vyos:"antenna-count,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtAntennaPatternFixed types.Bool   `tfsdk:"antenna_pattern_fixed" vyos:"antenna-pattern-fixed,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtBeamform            types.List   `tfsdk:"beamform" vyos:"beamform,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtChannelSetWIDth     types.String `tfsdk:"channel_set_width" vyos:"channel-set-width,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtLdpc                types.Bool   `tfsdk:"ldpc" vyos:"ldpc,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtLinkAdaptation      types.String `tfsdk:"link_adaptation" vyos:"link-adaptation,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtMaxMpduExp          types.Number `tfsdk:"max_mpdu_exp" vyos:"max-mpdu-exp,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtMaxMpdu             types.String `tfsdk:"max_mpdu" vyos:"max-mpdu,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtShortGi             types.List   `tfsdk:"short_gi" vyos:"short-gi,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtTxPowersave         types.Bool   `tfsdk:"tx_powersave" vyos:"tx-powersave,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtVhtCf               types.Bool   `tfsdk:"vht_cf" vyos:"vht-cf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesWirelessCapabilitiesVhtCenterChannelFreq *InterfacesWirelessCapabilitiesVhtCenterChannelFreq `tfsdk:"center_channel_freq" vyos:"center-channel-freq,omitempty"`
	NodeInterfacesWirelessCapabilitiesVhtStbc              *InterfacesWirelessCapabilitiesVhtStbc              `tfsdk:"stbc" vyos:"stbc,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessCapabilitiesVht) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"antenna_count": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of antennas on this card

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-8  &emsp; |  Number of antennas for this card  |

`,
		},

		"antenna_pattern_fixed": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set if antenna pattern does not change during the lifetime of an association

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"beamform": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Beamforming capabilities

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  single-user-beamformer  &emsp; |  Support for operation as single user beamformer  |
    |  single-user-beamformee  &emsp; |  Support for operation as single user beamformee  |
    |  multi-user-beamformer  &emsp; |  Support for operation as multi user beamformer  |
    |  multi-user-beamformee  &emsp; |  Support for operation as multi user beamformee  |

`,
		},

		"channel_set_width": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VHT operating Channel width

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  0  &emsp; |  20 or 40 MHz channel width  |
    |  1  &emsp; |  80 MHz channel width  |
    |  2  &emsp; |  160 MHz channel width  |
    |  3  &emsp; |  80+80 MHz channel width  |

`,
		},

		"ldpc": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable LDPC (Low Density Parity Check) coding capability

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"link_adaptation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VHT link adaptation capabilities

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  unsolicited  &emsp; |  Station provides only unsolicited VHT MFB  |
    |  both  &emsp; |  Station can provide VHT MFB in response to VHT MRQ and unsolicited VHT MFB  |

`,
		},

		"max_mpdu_exp": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set the maximum length of A-MPDU pre-EOF padding that the station can receive

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-7  &emsp; |  Maximum length of A-MPDU pre-EOF padding = 2 pow(13 + x) -1 octets  |

`,
		},

		"max_mpdu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Increase Maximum MPDU length to 7991 or 11454 octets (otherwise: 3895 octets)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  7991  &emsp; |  ncrease Maximum MPDU length to 7991 octets  |
    |  11454  &emsp; |  ncrease Maximum MPDU length to 11454 octets  |

`,
		},

		"short_gi": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Short GI capabilities

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  80  &emsp; |  Short GI for 80 MHz  |
    |  160  &emsp; |  Short GI for 160 MHz  |

`,
		},

		"tx_powersave": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable VHT TXOP Power Save Mode

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vht_cf": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Station supports receiving VHT variant HT Control field

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"center_channel_freq": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesVhtCenterChannelFreq{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `VHT operating channel center frequency

`,
		},

		"stbc": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesVhtStbc{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Support for sending and receiving PPDU using STBC (Space Time Block Coding)

`,
		},
	}
}
