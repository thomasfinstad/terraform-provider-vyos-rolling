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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (allowas-in) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasInNumber types.Number `tfsdk:"number" vyos:"number,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (number) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of occurrences of AS number

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  1-10    |  Number of times AS is allowed in path  |
`,
			Description: `Number of occurrences of AS number

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  1-10    |  Number of times AS is allowed in path  |
`,
		},

		// TagNodes

		// Nodes

	}
}
