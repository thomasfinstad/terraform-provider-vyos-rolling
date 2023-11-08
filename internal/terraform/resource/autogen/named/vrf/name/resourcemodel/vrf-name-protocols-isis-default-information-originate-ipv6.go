// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsIsisDefaultInformationOriginateIPvsix describes the resource data model.
type VrfNameProtocolsIsisDefaultInformationOriginateIPvsix struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne *VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo *VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisDefaultInformationOriginateIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Distribute default route into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Distribute default route into level-2

`,
		},
	}
}
