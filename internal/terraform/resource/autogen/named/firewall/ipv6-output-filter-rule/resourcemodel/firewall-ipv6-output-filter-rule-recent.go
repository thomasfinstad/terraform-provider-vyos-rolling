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

var _ helpers.VyosResourceDataModel = &FirewallIPvsixOutputFilterRuleRecent{}

// FirewallIPvsixOutputFilterRuleRecent describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixOutputFilterRuleRecent struct {
	// LeafNodes
	LeafFirewallIPvsixOutputFilterRuleRecentCount types.Number `tfsdk:"count" vyos:"count,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleRecentTime  types.String `tfsdk:"time" vyos:"time,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputFilterRuleRecent) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen in the last second/minute/hour

    |  Format  |  Description                                           |
    |----------|--------------------------------------------------------|
    |  second  |  Source addresses seen COUNT times in the last second  |
    |  minute  |  Source addresses seen COUNT times in the last minute  |
    |  hour    |  Source addresses seen COUNT times in the last hour    |
`,
			Description: `Source addresses seen in the last second/minute/hour

    |  Format  |  Description                                           |
    |----------|--------------------------------------------------------|
    |  second  |  Source addresses seen COUNT times in the last second  |
    |  minute  |  Source addresses seen COUNT times in the last minute  |
    |  hour    |  Source addresses seen COUNT times in the last hour    |
`,
		},

		// TagNodes

		// Nodes

	}
}
