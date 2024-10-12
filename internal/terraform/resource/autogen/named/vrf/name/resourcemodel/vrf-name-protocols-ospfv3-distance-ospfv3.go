// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfvthreeDistanceOspfvthree{}

// VrfNameProtocolsOspfvthreeDistanceOspfvthree describes the resource data model.
type VrfNameProtocolsOspfvthreeDistanceOspfvthree struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeExternal  types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeInterArea types.Number `tfsdk:"inter_area" vyos:"inter-area,omitempty"`
	LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeIntraArea types.Number `tfsdk:"intra_area" vyos:"intra-area,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeDistanceOspfvthree) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for external routes

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  1-255   |  Distance for external routes  |
`,
			Description: `Distance for external routes

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  1-255   |  Distance for external routes  |
`,
		},

		"inter_area": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for inter-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for inter-area routes  |
`,
			Description: `Distance for inter-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for inter-area routes  |
`,
		},

		"intra_area": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for intra-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for intra-area routes  |
`,
			Description: `Distance for intra-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for intra-area routes  |
`,
		},

		// Nodes

	}
}
