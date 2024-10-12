// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapability{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapability describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapability struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrf *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrf `tfsdk:"orf" vyos:"orf,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapability) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"orf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise ORF capability to this peer

`,
			Description: `Advertise ORF capability to this peer

`,
		},
	}
}
