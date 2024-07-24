// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisRedistributeIPvfourOspf{}

// VrfNameProtocolsIsisRedistributeIPvfourOspf describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvfourOspf struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvfourOspfLevelOne *VrfNameProtocolsIsisRedistributeIPvfourOspfLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourOspfLevelTwo *VrfNameProtocolsIsisRedistributeIPvfourOspfLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfourOspf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourOspfLevelOne{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
			Description: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourOspfLevelTwo{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
			Description: `Redistribute into level-2

`,
		},
	}
}
