/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesWirelessCapabilitiesHe{}

// InterfacesWirelessCapabilitiesHe describes the resource data model.
type InterfacesWirelessCapabilitiesHe struct {
	// LeafNodes
	LeafInterfacesWirelessCapabilitiesHeChannelSetWIDth     types.String `tfsdk:"channel_set_width" vyos:"channel-set-width,omitempty"`
	LeafInterfacesWirelessCapabilitiesHeAntennaPatternFixed types.Bool   `tfsdk:"antenna_pattern_fixed" vyos:"antenna-pattern-fixed,omitempty"`
	LeafInterfacesWirelessCapabilitiesHeBssColor            types.String `tfsdk:"bss_color" vyos:"bss-color,omitempty"`
	LeafInterfacesWirelessCapabilitiesHeCodingScheme        types.Number `tfsdk:"coding_scheme" vyos:"coding-scheme,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesWirelessCapabilitiesHeCenterChannelFreq *InterfacesWirelessCapabilitiesHeCenterChannelFreq `tfsdk:"center_channel_freq" vyos:"center-channel-freq,omitempty"`
	NodeInterfacesWirelessCapabilitiesHeBeamform          *InterfacesWirelessCapabilitiesHeBeamform          `tfsdk:"beamform" vyos:"beamform,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessCapabilitiesHe) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"channel_set_width":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `HE operating channel width

    |  Format  |  Description                                                                  |
    |----------|-------------------------------------------------------------------------------|
    |  81      |  2.4GHz, 20 MHz channel width                                                 |
    |  83      |  2.4GHz, 40 MHz channel width, secondary 20MHz channel above primary channel  |
    |  84      |  2.4GHz, 40 MHz channel width, secondary 20MHz channel below primary channel  |
    |  131     |  6GHz, 20 MHz channel width                                                   |
    |  132     |  6GHz, 40 MHz channel width                                                   |
    |  133     |  6GHz, 80 MHz channel width                                                   |
    |  134     |  6GHz, 160 MHz channel width                                                  |
    |  135     |  6GHz, 80+80 MHz channel width                                                |
`,
			Description: `HE operating channel width

    |  Format  |  Description                                                                  |
    |----------|-------------------------------------------------------------------------------|
    |  81      |  2.4GHz, 20 MHz channel width                                                 |
    |  83      |  2.4GHz, 40 MHz channel width, secondary 20MHz channel above primary channel  |
    |  84      |  2.4GHz, 40 MHz channel width, secondary 20MHz channel below primary channel  |
    |  131     |  6GHz, 20 MHz channel width                                                   |
    |  132     |  6GHz, 40 MHz channel width                                                   |
    |  133     |  6GHz, 80 MHz channel width                                                   |
    |  134     |  6GHz, 160 MHz channel width                                                  |
    |  135     |  6GHz, 80+80 MHz channel width                                                |
`,
		},

		"antenna_pattern_fixed":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Tell the AP that antenna positions are fixed and will not change during the lifetime of an association

`,
			Description: `Tell the AP that antenna positions are fixed and will not change during the lifetime of an association

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"bss_color":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BSS coloring helps to prevent channel jamming when multiple APs use the same channels

`,
			Description: `BSS coloring helps to prevent channel jamming when multiple APs use the same channels

`,
		},

		"coding_scheme":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Spacial Stream and Modulation Coding Scheme settings

    |  Format  |  Description              |
    |----------|---------------------------|
    |  0       |  HE-MCS 0-7               |
    |  1       |  HE-MCS 0-9               |
    |  2       |  HE-MCS 0-11              |
    |  3       |  HE-MCS is not supported  |
`,
			Description: `Spacial Stream and Modulation Coding Scheme settings

    |  Format  |  Description              |
    |----------|---------------------------|
    |  0       |  HE-MCS 0-7               |
    |  1       |  HE-MCS 0-9               |
    |  2       |  HE-MCS 0-11              |
    |  3       |  HE-MCS is not supported  |
`,
		},

		// Nodes

		"center_channel_freq": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesHeCenterChannelFreq{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `HE operating channel center frequency

`,
			Description: `HE operating channel center frequency

`,
		},

		"beamform": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesHeBeamform{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `HE beamforming capabilities

`,
			Description: `HE beamforming capabilities

`,
		},
	}
}
