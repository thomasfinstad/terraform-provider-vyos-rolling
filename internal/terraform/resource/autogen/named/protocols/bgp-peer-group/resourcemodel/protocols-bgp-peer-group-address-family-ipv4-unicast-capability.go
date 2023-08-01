// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf `tfsdk:"orf" vyos:"orf,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"orf": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise ORF capability to this peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability) UnmarshalJSON(_ []byte) error {
	return nil
}