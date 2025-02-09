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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (as-path) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSetAsPath{}

// PolicyRouteMapRuleSetAsPath describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyRouteMapRuleSetAsPath struct {
	// LeafNodes
	LeafPolicyRouteMapRuleSetAsPathExclude       types.String `tfsdk:"exclude" vyos:"exclude,omitempty"`
	LeafPolicyRouteMapRuleSetAsPathPrepend       types.Number `tfsdk:"prepend" vyos:"prepend,omitempty"`
	LeafPolicyRouteMapRuleSetAsPathPrependLastAs types.Number `tfsdk:"prepend_last_as" vyos:"prepend-last-as,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetAsPath) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"exclude":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (exclude) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remove/exclude from the as-path attribute

    |  Format        |  Description                              |
    |----------------|-------------------------------------------|
    |  1-4294967295  |  AS number                                |
    |  all           |  Exclude all AS numbers from the as-path  |
`,
			Description: `Remove/exclude from the as-path attribute

    |  Format        |  Description                              |
    |----------------|-------------------------------------------|
    |  1-4294967295  |  AS number                                |
    |  all           |  Exclude all AS numbers from the as-path  |
`,
		},

		"prepend":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (prepend) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prepend to the as-path

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-4294967295  |  AS number    |
`,
			Description: `Prepend to the as-path

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-4294967295  |  AS number    |
`,
		},

		"prepend_last_as":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (prepend-last-as) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Use the last AS-number in the as-path

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-10    |  Number of times to insert  |
`,
			Description: `Use the last AS-number in the as-path

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-10    |  Number of times to insert  |
`,
		},

		// TagNodes

		// Nodes

	}
}
