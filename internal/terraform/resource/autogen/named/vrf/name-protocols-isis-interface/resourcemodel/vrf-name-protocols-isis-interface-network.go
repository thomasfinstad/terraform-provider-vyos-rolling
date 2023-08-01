// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsIsisInterfaceNetwork describes the resource data model.
type VrfNameProtocolsIsisInterfaceNetwork struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisInterfaceNetworkPointToPoint types.Bool `tfsdk:"point_to_point" vyos:"point-to-point,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterfaceNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"point_to_point": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `point-to-point network type

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisInterfaceNetwork) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisInterfaceNetwork) UnmarshalJSON(_ []byte) error {
	return nil
}