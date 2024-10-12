// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsRIPngDistributeListInterfaceAccessList{}

// ProtocolsRIPngDistributeListInterfaceAccessList describes the resource data model.
type ProtocolsRIPngDistributeListInterfaceAccessList struct {
	// LeafNodes
	LeafProtocolsRIPngDistributeListInterfaceAccessListIn  types.Number `tfsdk:"in" vyos:"in,omitempty"`
	LeafProtocolsRIPngDistributeListInterfaceAccessListOut types.Number `tfsdk:"out" vyos:"out,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngDistributeListInterfaceAccessList) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"in": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access list to apply to input packets

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  u32     |  Access list to apply to input packets  |
`,
			Description: `Access list to apply to input packets

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  u32     |  Access list to apply to input packets  |
`,
		},

		"out": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access list to apply to output packets

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  u32     |  Access list to apply to output packets  |
`,
			Description: `Access list to apply to output packets

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  u32     |  Access list to apply to output packets  |
`,
		},

		// Nodes

	}
}
