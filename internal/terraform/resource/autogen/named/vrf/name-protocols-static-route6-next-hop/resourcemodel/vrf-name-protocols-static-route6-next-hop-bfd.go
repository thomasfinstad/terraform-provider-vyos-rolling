// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsStaticRoutesixNextHopBfd describes the resource data model.
type VrfNameProtocolsStaticRoutesixNextHopBfd struct {
	// LeafNodes
	LeafVrfNameProtocolsStaticRoutesixNextHopBfdProfile types.String `tfsdk:"profile" vyos:"profile,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsStaticRoutesixNextHopBfdMultiHop *VrfNameProtocolsStaticRoutesixNextHopBfdMultiHop `tfsdk:"multi_hop" vyos:"multi-hop,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRoutesixNextHopBfd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  BFD profile name  |

`,
		},

		// Nodes

		"multi_hop": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsStaticRoutesixNextHopBfdMultiHop{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use BFD multi hop session

`,
		},
	}
}
