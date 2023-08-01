// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// InterfacesWirelessSecURIty describes the resource data model.
type InterfacesWirelessSecURIty struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesWirelessSecURItyWep *InterfacesWirelessSecURItyWep `tfsdk:"wep" vyos:"wep,omitempty"`
	NodeInterfacesWirelessSecURItyWpa *InterfacesWirelessSecURItyWpa `tfsdk:"wpa" vyos:"wpa,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessSecURIty) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"wep": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURItyWep{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Wired Equivalent Privacy (WEP) parameters

`,
		},

		"wpa": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURItyWpa{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Wifi Protected Access (WPA) parameters

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessSecURIty) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWirelessSecURIty) UnmarshalJSON(_ []byte) error {
	return nil
}