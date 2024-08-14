// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesWirelessSecURItyWpa{}

// InterfacesWirelessSecURItyWpa describes the resource data model.
type InterfacesWirelessSecURItyWpa struct {
	// LeafNodes
	LeafInterfacesWirelessSecURItyWpaCIPher          types.List   `tfsdk:"cipher" vyos:"cipher,omitempty"`
	LeafInterfacesWirelessSecURItyWpaGroupCIPher     types.String `tfsdk:"group_cipher" vyos:"group-cipher,omitempty"`
	LeafInterfacesWirelessSecURItyWpaGroupMgmtCIPher types.String `tfsdk:"group_mgmt_cipher" vyos:"group-mgmt-cipher,omitempty"`
	LeafInterfacesWirelessSecURItyWpaMode            types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafInterfacesWirelessSecURItyWpaUsername        types.String `tfsdk:"username" vyos:"username,omitempty"`
	LeafInterfacesWirelessSecURItyWpaPassphrase      types.String `tfsdk:"passphrase" vyos:"passphrase,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesWirelessSecURItyWpaRadius *InterfacesWirelessSecURItyWpaRadius `tfsdk:"radius" vyos:"radius,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessSecURItyWpa) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cipher": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Cipher suite for WPA unicast packets

    |  Format    |  Description                                                                                 |
    |------------|----------------------------------------------------------------------------------------------|
    |  GCMP-256  |  AES in Galois/counter mode with 256-bit key                                                 |
    |  GCMP      |  AES in Galois/counter mode with 128-bit key                                                 |
    |  CCMP-256  |  AES in Counter mode with CBC-MAC with 256-bit key                                           |
    |  CCMP      |  AES in Counter mode with CBC-MAC [RFC 3610, IEEE 802.11i/D7.0] (supported on all WPA2 APs)  |
    |  TKIP      |  Temporal Key Integrity Protocol [IEEE 802.11i/D7.0]                                         |
`,
			Description: `Cipher suite for WPA unicast packets

    |  Format    |  Description                                                                                 |
    |------------|----------------------------------------------------------------------------------------------|
    |  GCMP-256  |  AES in Galois/counter mode with 256-bit key                                                 |
    |  GCMP      |  AES in Galois/counter mode with 128-bit key                                                 |
    |  CCMP-256  |  AES in Counter mode with CBC-MAC with 256-bit key                                           |
    |  CCMP      |  AES in Counter mode with CBC-MAC [RFC 3610, IEEE 802.11i/D7.0] (supported on all WPA2 APs)  |
    |  TKIP      |  Temporal Key Integrity Protocol [IEEE 802.11i/D7.0]                                         |
`,
		},

		"group_cipher": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Cipher suite for WPA multicast and broadcast packets

    |  Format    |  Description                                                                                 |
    |------------|----------------------------------------------------------------------------------------------|
    |  GCMP-256  |  AES in Galois/counter mode with 256-bit key                                                 |
    |  GCMP      |  AES in Galois/counter mode with 128-bit key                                                 |
    |  CCMP-256  |  AES in Counter mode with CBC-MAC with 256-bit key                                           |
    |  CCMP      |  AES in Counter mode with CBC-MAC [RFC 3610, IEEE 802.11i/D7.0] (supported on all WPA2 APs)  |
    |  TKIP      |  Temporal Key Integrity Protocol [IEEE 802.11i/D7.0]                                         |
`,
			Description: `Cipher suite for WPA multicast and broadcast packets

    |  Format    |  Description                                                                                 |
    |------------|----------------------------------------------------------------------------------------------|
    |  GCMP-256  |  AES in Galois/counter mode with 256-bit key                                                 |
    |  GCMP      |  AES in Galois/counter mode with 128-bit key                                                 |
    |  CCMP-256  |  AES in Counter mode with CBC-MAC with 256-bit key                                           |
    |  CCMP      |  AES in Counter mode with CBC-MAC [RFC 3610, IEEE 802.11i/D7.0] (supported on all WPA2 APs)  |
    |  TKIP      |  Temporal Key Integrity Protocol [IEEE 802.11i/D7.0]                                         |
`,
		},

		"group_mgmt_cipher": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group management cipher suite. All the stations connecting to the BSS will also need to support the selected cipher

`,
			Description: `Group management cipher suite. All the stations connecting to the BSS will also need to support the selected cipher

`,

			// Default:          stringdefault.StaticString(`AES-128-CMAC`),
			Computed: true,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `WPA mode

    |  Format    |  Description                                                                        |
    |------------|-------------------------------------------------------------------------------------|
    |  wpa       |  WPA (IEEE 802.11i/D3.0)                                                            |
    |  wpa2      |  WPA2 (full IEEE 802.11i/RSN)                                                       |
    |  wpa+wpa2  |  Allow both WPA and WPA2                                                            |
    |  wpa3      |  WPA3 (required for 802.11ax, you must also set mgmt-frame-protection as required)  |
`,
			Description: `WPA mode

    |  Format    |  Description                                                                        |
    |------------|-------------------------------------------------------------------------------------|
    |  wpa       |  WPA (IEEE 802.11i/D3.0)                                                            |
    |  wpa2      |  WPA2 (full IEEE 802.11i/RSN)                                                       |
    |  wpa+wpa2  |  Allow both WPA and WPA2                                                            |
    |  wpa3      |  WPA3 (required for 802.11ax, you must also set mgmt-frame-protection as required)  |
`,

			// Default:          stringdefault.StaticString(`wpa+wpa2`),
			Computed: true,
		},

		"username": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Username     |
`,
			Description: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Username     |
`,
		},

		"passphrase": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `WPA passphrase. If you are using special characters in the WPA passphrase then single quotes are required.

    |  Format  |  Description                                                                                                                |
    |----------|-----------------------------------------------------------------------------------------------------------------------------|
    |  txt     |  Passphrase of at least 8 but not more than 63 printable characters for WPA-Personal and any passphrase for WPA-Enterprise  |
`,
			Description: `WPA passphrase. If you are using special characters in the WPA passphrase then single quotes are required.

    |  Format  |  Description                                                                                                                |
    |----------|-----------------------------------------------------------------------------------------------------------------------------|
    |  txt     |  Passphrase of at least 8 but not more than 63 printable characters for WPA-Personal and any passphrase for WPA-Enterprise  |
`,
		},

		// Nodes

		"radius": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURItyWpaRadius{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `RADIUS based user authentication

`,
			Description: `RADIUS based user authentication

`,
		},
	}
}