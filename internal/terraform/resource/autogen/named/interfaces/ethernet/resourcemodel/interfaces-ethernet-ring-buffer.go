// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesEthernetRingBuffer{}

// InterfacesEthernetRingBuffer describes the resource data model.
type InterfacesEthernetRingBuffer struct {
	// LeafNodes
	LeafInterfacesEthernetRingBufferRx types.Number `tfsdk:"rx" vyos:"rx,omitempty"`
	LeafInterfacesEthernetRingBufferTx types.Number `tfsdk:"tx" vyos:"tx,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetRingBuffer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"rx": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `RX ring buffer

    |  Format    |  Description       |
    |------------|--------------------|
    |  80-16384  |  ring buffer size  |
`,
			Description: `RX ring buffer

    |  Format    |  Description       |
    |------------|--------------------|
    |  80-16384  |  ring buffer size  |
`,
		},

		"tx": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TX ring buffer

    |  Format    |  Description       |
    |------------|--------------------|
    |  80-16384  |  ring buffer size  |
`,
			Description: `TX ring buffer

    |  Format    |  Description       |
    |------------|--------------------|
    |  80-16384  |  ring buffer size  |
`,
		},

		// Nodes

	}
}