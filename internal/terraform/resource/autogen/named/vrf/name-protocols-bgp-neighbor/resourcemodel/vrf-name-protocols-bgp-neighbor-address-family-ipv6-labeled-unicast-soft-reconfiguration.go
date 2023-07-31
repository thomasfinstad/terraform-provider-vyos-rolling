// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfigurationInbound types.Bool `tfsdk:"inbound" vyos:"inbound,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"inbound": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable inbound soft reconfiguration

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastSoftReconfiguration) UnmarshalJSON(_ []byte) error {
	return nil
}
