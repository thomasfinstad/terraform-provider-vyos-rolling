// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList *ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
		},
	}
}
