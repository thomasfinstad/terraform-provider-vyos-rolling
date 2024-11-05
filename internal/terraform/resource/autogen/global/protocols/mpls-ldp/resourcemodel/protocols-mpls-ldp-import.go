/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsMplsLdpImport{}

// ProtocolsMplsLdpImport describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsMplsLdpImport struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeProtocolsMplsLdpImportIPvfour *ProtocolsMplsLdpImportIPvfour `tfsdk:"ipv4" vyos:"ipv4,omitempty"`

	NodeProtocolsMplsLdpImportIPvsix *ProtocolsMplsLdpImportIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpImport) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"ipv4": schema.SingleNestedAttribute{
			Attributes: ProtocolsMplsLdpImportIPvfour{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 parameters

`,
			Description: `IPv4 parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: ProtocolsMplsLdpImportIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 parameters

`,
			Description: `IPv6 parameters

`,
		},
	}
}
