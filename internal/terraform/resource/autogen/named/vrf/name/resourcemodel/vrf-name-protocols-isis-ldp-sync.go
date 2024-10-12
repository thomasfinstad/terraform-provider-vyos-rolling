// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisLdpSync{}

// VrfNameProtocolsIsisLdpSync describes the resource data model.
type VrfNameProtocolsIsisLdpSync struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisLdpSyncHolddown types.Number `tfsdk:"holddown" vyos:"holddown,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisLdpSync) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"holddown": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hold down timer for LDP-IGP cost restoration

    |  Format   |  Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    |  0-10000  |  Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |
`,
			Description: `Hold down timer for LDP-IGP cost restoration

    |  Format   |  Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    |  0-10000  |  Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |
`,
		},

		// Nodes

	}
}
