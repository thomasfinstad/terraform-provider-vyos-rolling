// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf{}

// ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList *ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
			Description: `Advertise prefix-list ORF capability to this peer

`,
		},
	}
}
