// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessSecURItyWpaRadius describes the resource data model.
type InterfacesWirelessSecURItyWpaRadius struct {
	// LeafNodes
	LeafInterfacesWirelessSecURItyWpaRadiusSourceAddress types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesWirelessSecURItyWpaRadiusServer bool `tfsdk:"server" vyos:"server,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessSecURItyWpaRadius) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 source address used to initiate connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 source address  |

`,
		},

		// Nodes

	}
}
