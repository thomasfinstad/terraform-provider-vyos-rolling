// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListExport types.String `tfsdk:"export" vyos:"export,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListImport types.String `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4-Prefix-list to filter outgoing route updates to this peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of IPv4 prefix-list  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4-Prefix-list to filter incoming route updates from this peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of IPv4 prefix-list  |

`,
		},

		// Nodes

	}
}
