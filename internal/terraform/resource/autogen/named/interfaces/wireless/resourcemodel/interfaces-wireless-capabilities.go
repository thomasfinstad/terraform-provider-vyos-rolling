// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessCapabilities describes the resource data model.
type InterfacesWirelessCapabilities struct {
	// LeafNodes
	LeafInterfacesWirelessCapabilitiesRequireHt  types.Bool `tfsdk:"require_ht" vyos:"require-ht,omitempty"`
	LeafInterfacesWirelessCapabilitiesRequireVht types.Bool `tfsdk:"require_vht" vyos:"require-vht,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesWirelessCapabilitiesHt  *InterfacesWirelessCapabilitiesHt  `tfsdk:"ht" vyos:"ht,omitempty"`
	NodeInterfacesWirelessCapabilitiesVht *InterfacesWirelessCapabilitiesVht `tfsdk:"vht" vyos:"vht,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessCapabilities) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"require_ht": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Require stations to support HT PHY (reject association if they do not)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"require_vht": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Require stations to support VHT PHY (reject association if they do not)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"ht": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesHt{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `HT (High Throughput) settings

`,
		},

		"vht": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessCapabilitiesVht{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `VHT (Very High Throughput) settings

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessCapabilities) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWirelessCapabilities) UnmarshalJSON(_ []byte) error {
	return nil
}