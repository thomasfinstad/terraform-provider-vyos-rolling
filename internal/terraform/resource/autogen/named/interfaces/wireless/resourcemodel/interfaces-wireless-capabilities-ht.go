/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (ht) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesWirelessCapabilitiesHt{}

// InterfacesWirelessCapabilitiesHt describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesWirelessCapabilitiesHt struct {
	// LeafNodes
	LeafInterfacesWirelessCapabilitiesHtFourzeromhzIncapable types.Bool   `tfsdk:"_40mhz_incapable" vyos:"40mhz-incapable,omitempty"`
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

	// TagNodes

	// Nodes

	NodeInterfacesWirelessCapabilitiesHtStbc *InterfacesWirelessCapabilitiesHtStbc `tfsdk:"stbc" vyos:"stbc,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessCapabilitiesHt) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"_40mhz_incapable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (40mhz-incapable) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `40MHz intolerance, use 20MHz only!

`,
			Description: `40MHz intolerance, use 20MHz only!

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"auto_powersave":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (auto-powersave) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable WMM-PS unscheduled automatic power save delivery [U-APSD]

`,
			Description: `Enable WMM-PS unscheduled automatic power save delivery [U-APSD]

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"channel_set_width":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (channel-set-width) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Supported channel set width

    |  Format  |  Description                                                                                      |
    |----------|---------------------------------------------------------------------------------------------------|
    |  ht20    |  Supported channel set width both 20 MHz only                                                     |
    |  ht40+   |  Supported channel set width both 20 MHz and 40 MHz with secondary channel above primary channel  |
    |  ht40-   |  Supported channel set width both 20 MHz and 40 MHz with secondary channel below primary channel  |
`,
			Description: `Supported channel set width

    |  Format  |  Description                                                                                      |
    |----------|---------------------------------------------------------------------------------------------------|
    |  ht20    |  Supported channel set width both 20 MHz only                                                     |
    |  ht40+   |  Supported channel set width both 20 MHz and 40 MHz with secondary channel above primary channel  |
    |  ht40-   |  Supported channel set width both 20 MHz and 40 MHz with secondary channel below primary channel  |
`,
		},

		"delayed_block_ack":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (delayed-block-ack) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable HT-delayed block ack

`,
			Description: `Enable HT-delayed block ack

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"dsss_cck_40":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (dsss-cck-40) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable DSSS_CCK-40

`,
			Description: `Enable DSSS_CCK-40

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"greenfield":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (greenfield) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable HT-greenfield

`,
			Description: `Enable HT-greenfield

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ldpc":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ldpc) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable LDPC coding capability

`,
			Description: `Enable LDPC coding capability

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"lsig_protection":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (lsig-protection) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable L-SIG TXOP protection capability

`,
			Description: `Enable L-SIG TXOP protection capability

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"max_amsdu":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (max-amsdu) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set maximum A-MSDU length

    |  Format  |  Description                               |
    |----------|--------------------------------------------|
    |  3839    |  Set maximum A-MSDU length to 3839 octets  |
    |  7935    |  Set maximum A-MSDU length to 7935 octets  |
`,
			Description: `Set maximum A-MSDU length

    |  Format  |  Description                               |
    |----------|--------------------------------------------|
    |  3839    |  Set maximum A-MSDU length to 3839 octets  |
    |  7935    |  Set maximum A-MSDU length to 7935 octets  |
`,
		},

		"short_gi":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (short-gi) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Short GI capabilities

    |  Format  |  Description          |
    |----------|-----------------------|
    |  20      |  Short GI for 20 MHz  |
    |  40      |  Short GI for 40 MHz  |
`,
			Description: `Short GI capabilities

    |  Format  |  Description          |
    |----------|-----------------------|
    |  20      |  Short GI for 20 MHz  |
    |  40      |  Short GI for 40 MHz  |
`,
		},

		"smps":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (smps) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Spatial Multiplexing Power Save (SMPS) settings

    |  Format   |  Description                                   |
    |-----------|------------------------------------------------|
    |  static   |  STATIC Spatial Multiplexing (SM) Power Save   |
    |  dynamic  |  DYNAMIC Spatial Multiplexing (SM) Power Save  |
`,
			Description: `Spatial Multiplexing Power Save (SMPS) settings

    |  Format   |  Description                                   |
    |-----------|------------------------------------------------|
    |  static   |  STATIC Spatial Multiplexing (SM) Power Save   |
    |  dynamic  |  DYNAMIC Spatial Multiplexing (SM) Power Save  |
`,
		},

		// TagNodes

		// Nodes

		"stbc": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesHtStbc{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Support for sending and receiving PPDU using STBC (Space Time Block Coding)

`,
			Description: `Support for sending and receiving PPDU using STBC (Space Time Block Coding)

`,
		},
	}
}
