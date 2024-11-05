/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpPeerGroupAddressFamilyIPvfourVpnPrefixList{}

// ProtocolsBgpPeerGroupAddressFamilyIPvfourVpnPrefixList describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpPeerGroupAddressFamilyIPvfourVpnPrefixList struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourVpnPrefixListExport types.String `tfsdk:"export" vyos:"export,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourVpnPrefixListImport types.String `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourVpnPrefixList) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4-Prefix-list to filter outgoing route updates to this peer

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv4 prefix-list  |
`,
			Description: `IPv4-Prefix-list to filter outgoing route updates to this peer

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv4 prefix-list  |
`,
		},

		"import":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4-Prefix-list to filter incoming route updates from this peer

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv4 prefix-list  |
`,
			Description: `IPv4-Prefix-list to filter incoming route updates from this peer

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv4 prefix-list  |
`,
		},

		// TagNodes

		// Nodes

	}
}
