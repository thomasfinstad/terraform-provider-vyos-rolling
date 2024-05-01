// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrf{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrf describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrf struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
			Description: `Advertise prefix-list ORF capability to this peer

`,
		},
	}
}
