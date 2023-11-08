// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsIsisRedistributeIPvfourStatic describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvfourStatic struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvfourStaticLevelOne *VrfNameProtocolsIsisRedistributeIPvfourStaticLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourStaticLevelTwo *VrfNameProtocolsIsisRedistributeIPvfourStaticLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfourStatic) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourStaticLevelOne{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourStaticLevelTwo{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
		},
	}
}
