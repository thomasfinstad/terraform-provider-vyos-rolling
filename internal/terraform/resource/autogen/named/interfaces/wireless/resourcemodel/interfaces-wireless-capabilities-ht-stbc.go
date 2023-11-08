// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessCapabilitiesHtStbc describes the resource data model.
type InterfacesWirelessCapabilitiesHtStbc struct {
	// LeafNodes
	LeafInterfacesWirelessCapabilitiesHtStbcRx types.String `tfsdk:"rx" vyos:"rx,omitempty"`
	LeafInterfacesWirelessCapabilitiesHtStbcTx types.Bool   `tfsdk:"tx" vyos:"tx,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessCapabilitiesHtStbc) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"rx": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable receiving PPDU using STBC (Space Time Block Coding)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  [1-3]+  &emsp; |  Number of spacial streams that can use RX STBC  |

`,
		},

		"tx": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable sending PPDU using STBC (Space Time Block Coding)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
