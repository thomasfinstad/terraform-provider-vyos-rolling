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

var _ helpers.VyosResourceDataModel = &PolicyLocalRouteRuleSet{}

// PolicyLocalRouteRuleSet describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyLocalRouteRuleSet struct {
	// LeafNodes
	LeafPolicyLocalRouteRuleSetTable types.Number `tfsdk:"table" vyos:"table,omitempty"`
	LeafPolicyLocalRouteRuleSetVrf   types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRouteRuleSet) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"table":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Routing table to forward packet with

    |  Format  |  Description   |
    |----------|----------------|
    |  1-200   |  Table number  |
`,
			Description: `Routing table to forward packet with

    |  Format  |  Description   |
    |----------|----------------|
    |  1-200   |  Table number  |
`,
		},

		"vrf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF to forward packet with

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  txt      |  VRF instance name                |
    |  default  |  Forward into default global VRF  |
`,
			Description: `VRF to forward packet with

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  txt      |  VRF instance name                |
    |  default  |  Forward into default global VRF  |
`,
		},

		// TagNodes

		// Nodes

	}
}
