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

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastDistributeList{}

// ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastDistributeList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastDistributeList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastDistributeListExport types.Number `tfsdk:"export" vyos:"export,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastDistributeListImport types.Number `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastDistributeList) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		// Nodes

	}
}
