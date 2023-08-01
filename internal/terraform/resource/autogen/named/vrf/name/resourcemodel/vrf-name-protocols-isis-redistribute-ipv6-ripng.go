// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsIsisRedistributeIPvsixRIPng describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvsixRIPng struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvsixRIPngLevelOne *VrfNameProtocolsIsisRedistributeIPvsixRIPngLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixRIPngLevelTwo *VrfNameProtocolsIsisRedistributeIPvsixRIPngLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixRIPng) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixRIPngLevelOne{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixRIPngLevelTwo{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisRedistributeIPvsixRIPng) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisRedistributeIPvsixRIPng) UnmarshalJSON(_ []byte) error {
	return nil
}