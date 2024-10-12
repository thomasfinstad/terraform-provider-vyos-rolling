// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisSegmentRoutingLocalBlock{}

// VrfNameProtocolsIsisSegmentRoutingLocalBlock describes the resource data model.
type VrfNameProtocolsIsisSegmentRoutingLocalBlock struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisSegmentRoutingLocalBlockLowLabelValue  types.Number `tfsdk:"low_label_value" vyos:"low-label-value,omitempty"`
	LeafVrfNameProtocolsIsisSegmentRoutingLocalBlockHighLabelValue types.Number `tfsdk:"high_label_value" vyos:"high-label-value,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisSegmentRoutingLocalBlock) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"low_label_value": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label lower bound

    |  Format      |  Description                                   |
    |--------------|------------------------------------------------|
    |  16-1048575  |  Label value (recommended minimum value: 300)  |
`,
			Description: `MPLS label lower bound

    |  Format      |  Description                                   |
    |--------------|------------------------------------------------|
    |  16-1048575  |  Label value (recommended minimum value: 300)  |
`,
		},

		"high_label_value": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label upper bound

    |  Format      |  Description  |
    |--------------|---------------|
    |  16-1048575  |  Label value  |
`,
			Description: `MPLS label upper bound

    |  Format      |  Description  |
    |--------------|---------------|
    |  16-1048575  |  Label value  |
`,
		},

		// Nodes

	}
}
