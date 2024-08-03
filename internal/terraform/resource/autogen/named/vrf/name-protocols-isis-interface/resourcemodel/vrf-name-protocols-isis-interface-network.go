// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisInterfaceNetwork{}

// VrfNameProtocolsIsisInterfaceNetwork describes the resource data model.
type VrfNameProtocolsIsisInterfaceNetwork struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisInterfaceNetworkPointToPoint types.Bool `tfsdk:"point_to_point" vyos:"point-to-point,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterfaceNetwork) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"point_to_point": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `point-to-point network type

`,
			Description: `point-to-point network type

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
