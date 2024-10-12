// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsStaticRouteNextHopBfd{}

// ProtocolsStaticRouteNextHopBfd describes the resource data model.
type ProtocolsStaticRouteNextHopBfd struct {
	// LeafNodes
	LeafProtocolsStaticRouteNextHopBfdProfile types.String `tfsdk:"profile" vyos:"profile,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsStaticRouteNextHopBfdMultiHop *ProtocolsStaticRouteNextHopBfdMultiHop `tfsdk:"multi_hop" vyos:"multi-hop,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticRouteNextHopBfd) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
			Description: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
		},

		// Nodes

		"multi_hop": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticRouteNextHopBfdMultiHop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Use BFD multi hop session

`,
			Description: `Use BFD multi hop session

`,
		},
	}
}
