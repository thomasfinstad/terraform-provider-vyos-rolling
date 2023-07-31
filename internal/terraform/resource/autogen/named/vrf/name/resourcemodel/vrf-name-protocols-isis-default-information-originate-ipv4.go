// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsIsisDefaultInformationOriginateIPvfour describes the resource data model.
type VrfNameProtocolsIsisDefaultInformationOriginateIPvfour struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOne *VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo *VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisDefaultInformationOriginateIPvfour) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOne{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Distribute default route into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Distribute default route into level-2

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisDefaultInformationOriginateIPvfour) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisDefaultInformationOriginateIPvfour) UnmarshalJSON(_ []byte) error {
	return nil
}
