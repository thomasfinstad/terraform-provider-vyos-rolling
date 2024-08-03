// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesBondingEvpn{}

// InterfacesBondingEvpn describes the resource data model.
type InterfacesBondingEvpn struct {
	// LeafNodes
	LeafInterfacesBondingEvpnEsDfPref types.Number `tfsdk:"es_df_pref" vyos:"es-df-pref,omitempty"`
	LeafInterfacesBondingEvpnEsID     types.String `tfsdk:"es_id" vyos:"es-id,omitempty"`
	LeafInterfacesBondingEvpnEsSysMac types.String `tfsdk:"es_sys_mac" vyos:"es-sys-mac,omitempty"`
	LeafInterfacesBondingEvpnUplink   types.Bool   `tfsdk:"uplink" vyos:"uplink,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingEvpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"es_df_pref": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Preference value used for designated forwarder (DF) election

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-65535  |  DF Preference value  |
`,
			Description: `Preference value used for designated forwarder (DF) election

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-65535  |  DF Preference value  |
`,
		},

		"es_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet segment identifier

    |  Format      |  Description                                 |
    |--------------|----------------------------------------------|
    |  1-16777215  |  Local discriminator                         |
    |  txt         |  10-byte ID - 00:11:22:33:44:55:AA:BB:CC:DD  |
`,
			Description: `Ethernet segment identifier

    |  Format      |  Description                                 |
    |--------------|----------------------------------------------|
    |  1-16777215  |  Local discriminator                         |
    |  txt         |  10-byte ID - 00:11:22:33:44:55:AA:BB:CC:DD  |
`,
		},

		"es_sys_mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet segment system MAC

    |  Format   |  Description  |
    |-----------|---------------|
    |  macaddr  |  MAC address  |
`,
			Description: `Ethernet segment system MAC

    |  Format   |  Description  |
    |-----------|---------------|
    |  macaddr  |  MAC address  |
`,
		},

		"uplink": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Uplink to the VXLAN core

`,
			Description: `Uplink to the VXLAN core

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
