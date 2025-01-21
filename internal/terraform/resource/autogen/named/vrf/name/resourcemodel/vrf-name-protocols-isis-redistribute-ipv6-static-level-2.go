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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (level-2) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisRedistributeIPvsixStaticLevelTwo{}

// VrfNameProtocolsIsisRedistributeIPvsixStaticLevelTwo describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsIsisRedistributeIPvsixStaticLevelTwo struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisRedistributeIPvsixStaticLevelTwoMetric   types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsIsisRedistributeIPvsixStaticLevelTwoRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixStaticLevelTwo) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (metric) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set default metric for circuit

    |  Format      |  Description           |
    |--------------|------------------------|
    |  0-16777215  |  Default metric value  |
`,
			Description: `Set default metric for circuit

    |  Format      |  Description           |
    |--------------|------------------------|
    |  0-16777215  |  Default metric value  |
`,
		},

		"route_map":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (route-map) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Specify route-map name to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
