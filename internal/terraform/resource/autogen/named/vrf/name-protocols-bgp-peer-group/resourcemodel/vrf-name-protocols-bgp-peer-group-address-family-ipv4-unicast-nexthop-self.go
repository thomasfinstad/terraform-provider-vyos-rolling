// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelfForce types.Bool `tfsdk:"force" vyos:"force,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"force": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set the next hop to self for reflected routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf) UnmarshalJSON(_ []byte) error {
	return nil
}
