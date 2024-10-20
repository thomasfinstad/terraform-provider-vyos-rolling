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

var _ helpers.VyosResourceDataModel = &FirewallIPvfourPreroutingRawRuleRecent{}

// FirewallIPvfourPreroutingRawRuleRecent describes the resource data model.
type FirewallIPvfourPreroutingRawRuleRecent struct {
	// LeafNodes
	LeafFirewallIPvfourPreroutingRawRuleRecentCount types.Number `tfsdk:"count" vyos:"count,omitempty"`
	LeafFirewallIPvfourPreroutingRawRuleRecentTime  types.String `tfsdk:"time" vyos:"time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourPreroutingRawRuleRecent) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"count":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		// Nodes

	}
}
