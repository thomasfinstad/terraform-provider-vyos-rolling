// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsEigrpMetric{}

// VrfNameProtocolsEigrpMetric describes the resource data model.
type VrfNameProtocolsEigrpMetric struct {
	// LeafNodes
	LeafVrfNameProtocolsEigrpMetricWeights types.Number `tfsdk:"weights" vyos:"weights,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsEigrpMetric) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"weights": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Modify metric coefficients

    |  Format  |  Description  |
    |----------|---------------|
    |  0-255   |  K1           |
`,
			Description: `Modify metric coefficients

    |  Format  |  Description  |
    |----------|---------------|
    |  0-255   |  K1           |
`,
		},

		// Nodes

	}
}
