// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpListen{}

// VrfNameProtocolsBgpListen describes the resource data model.
type VrfNameProtocolsBgpListen struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpListenLimit types.Number `tfsdk:"limit" vyos:"limit,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpListenRange bool `tfsdk:"range" vyos:"range,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpListen) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of dynamic neighbors that can be created

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-5000  |  BGP neighbor limit  |
`,
			Description: `Maximum number of dynamic neighbors that can be created

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-5000  |  BGP neighbor limit  |
`,
		},

		// Nodes

	}
}
