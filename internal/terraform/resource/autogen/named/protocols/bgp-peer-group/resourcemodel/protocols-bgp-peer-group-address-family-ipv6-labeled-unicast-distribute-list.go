/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (distribute-list) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDistributeList{}

// ProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDistributeList describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDistributeList struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDistributeListExport types.Number `tfsdk:"export" vyos:"export,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDistributeListImport types.Number `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDistributeList) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (export) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

    |  Format   |  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter outgoing route updates to this peer-group  |
`,
			Description: `Access-list to filter outgoing route updates to this peer-group

    |  Format   |  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter outgoing route updates to this peer-group  |
`,
		},

		"import":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (import) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

    |  Format   |  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter incoming route updates from this peer-group  |
`,
			Description: `Access-list to filter incoming route updates from this peer-group

    |  Format   |  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter incoming route updates from this peer-group  |
`,
		},

		// TagNodes

		// Nodes

	}
}
