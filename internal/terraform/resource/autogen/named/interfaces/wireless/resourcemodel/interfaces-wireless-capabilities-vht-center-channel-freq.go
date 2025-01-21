/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (center-channel-freq) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesWirelessCapabilitiesVhtCenterChannelFreq{}

// InterfacesWirelessCapabilitiesVhtCenterChannelFreq describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesWirelessCapabilitiesVhtCenterChannelFreq struct {
	// LeafNodes
	LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqOne types.Number `tfsdk:"freq_1" vyos:"freq-1,omitempty"`
	LeafInterfacesWirelessCapabilitiesVhtCenterChannelFreqFreqTwo types.Number `tfsdk:"freq_2" vyos:"freq-2,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessCapabilitiesVhtCenterChannelFreq) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"freq_1":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (freq-1) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `VHT operating channel center frequency - center freq 1 (for use with 80, 80+80 and 160 modes)

    |  Format  |  Description                                                                          |
    |----------|---------------------------------------------------------------------------------------|
    |  34-177  |  5Ghz (802.11 a/h/j/n/ac) center channel index (use 42 for primary 80MHz channel 36)  |
`,
			Description: `VHT operating channel center frequency - center freq 1 (for use with 80, 80+80 and 160 modes)

    |  Format  |  Description                                                                          |
    |----------|---------------------------------------------------------------------------------------|
    |  34-177  |  5Ghz (802.11 a/h/j/n/ac) center channel index (use 42 for primary 80MHz channel 36)  |
`,
		},

		"freq_2":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (freq-2) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `VHT operating channel center frequency - center freq 2 (for use with the 80+80 mode)

    |  Format  |  Description                                                                    |
    |----------|---------------------------------------------------------------------------------|
    |  34-177  |  5Ghz (802.11 ac) center channel index (use 58 for secondary 80MHz channel 52)  |
`,
			Description: `VHT operating channel center frequency - center freq 2 (for use with the 80+80 mode)

    |  Format  |  Description                                                                    |
    |----------|---------------------------------------------------------------------------------|
    |  34-177  |  5Ghz (802.11 ac) center channel index (use 58 for secondary 80MHz channel 52)  |
`,
		},

		// TagNodes

		// Nodes

	}
}
