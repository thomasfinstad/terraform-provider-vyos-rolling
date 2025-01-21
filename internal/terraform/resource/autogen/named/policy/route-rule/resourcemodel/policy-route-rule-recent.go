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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (recent) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteRuleRecent{}

// PolicyRouteRuleRecent describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyRouteRuleRecent struct {
	// LeafNodes
	LeafPolicyRouteRuleRecentCount types.Number `tfsdk:"count" vyos:"count,omitempty"`
	LeafPolicyRouteRuleRecentTime  types.Number `tfsdk:"time" vyos:"time,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleRecent) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"count":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (count) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen more than N times

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  1-255   |  Source addresses seen more than N times  |
`,
			Description: `Source addresses seen more than N times

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  1-255   |  Source addresses seen more than N times  |
`,
		},

		"time":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (time) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen in the last N seconds

    |  Format        |  Description                                  |
    |----------------|-----------------------------------------------|
    |  0-4294967295  |  Source addresses seen in the last N seconds  |
`,
			Description: `Source addresses seen in the last N seconds

    |  Format        |  Description                                  |
    |----------------|-----------------------------------------------|
    |  0-4294967295  |  Source addresses seen in the last N seconds  |
`,
		},

		// TagNodes

		// Nodes

	}
}
