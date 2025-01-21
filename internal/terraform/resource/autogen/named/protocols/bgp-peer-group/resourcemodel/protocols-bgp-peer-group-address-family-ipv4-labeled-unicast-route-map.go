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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (route-map) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMap{}

// ProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMap describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMap struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMapExport types.String `tfsdk:"export" vyos:"export,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMapImport types.String `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMap) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (export) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter outgoing route updates

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Route-map to filter outgoing route updates

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		"import":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (import) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter incoming route updates

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Route-map to filter incoming route updates

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
