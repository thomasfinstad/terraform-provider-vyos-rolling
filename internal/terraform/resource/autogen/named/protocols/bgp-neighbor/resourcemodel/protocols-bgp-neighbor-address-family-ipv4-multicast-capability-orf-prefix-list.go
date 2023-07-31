// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListReceive types.Bool `tfsdk:"receive" vyos:"receive,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListSend    types.Bool `tfsdk:"send" vyos:"send,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"receive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Capability to receive the ORF

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"send": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Capability to send the ORF

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList) UnmarshalJSON(_ []byte) error {
	return nil
}
