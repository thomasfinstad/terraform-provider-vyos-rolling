/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (orf) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrf{}

// ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrf describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrf struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList *ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
			Description: `Advertise prefix-list ORF capability to this peer

`,
		},
	}
}
