// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisDefaultInformationOriginateIPvsix{}

// VrfNameProtocolsIsisDefaultInformationOriginateIPvsix describes the resource data model.
type VrfNameProtocolsIsisDefaultInformationOriginateIPvsix struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne *VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo *VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisDefaultInformationOriginateIPvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Distribute default route into level-1

`,
			Description: `Distribute default route into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Distribute default route into level-2

`,
			Description: `Distribute default route into level-2

`,
		},
	}
}
