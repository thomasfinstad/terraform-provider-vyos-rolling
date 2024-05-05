// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisRedistributeIPvsixBabel{}

// VrfNameProtocolsIsisRedistributeIPvsixBabel describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvsixBabel struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne *VrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo *VrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixBabel) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
			Description: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
			Description: `Redistribute into level-2

`,
		},
	}
}
