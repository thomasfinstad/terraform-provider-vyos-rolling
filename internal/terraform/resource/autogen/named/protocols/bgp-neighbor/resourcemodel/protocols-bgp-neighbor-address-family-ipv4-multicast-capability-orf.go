// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrf describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrf struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrf) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrf) UnmarshalJSON(_ []byte) error {
	return nil
}