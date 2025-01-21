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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (limit) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRoutesixRuleLimit{}

// PolicyRoutesixRuleLimit describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyRoutesixRuleLimit struct {
	// LeafNodes
	LeafPolicyRoutesixRuleLimitBurst types.Number `tfsdk:"burst" vyos:"burst,omitempty"`
	LeafPolicyRoutesixRuleLimitRate  types.String `tfsdk:"rate" vyos:"rate,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleLimit) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"burst":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (burst) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of packets to allow in excess of rate

    |  Format        |  Description                                           |
    |----------------|--------------------------------------------------------|
    |  0-4294967295  |  Maximum number of packets to allow in excess of rate  |
`,
			Description: `Maximum number of packets to allow in excess of rate

    |  Format        |  Description                                           |
    |----------------|--------------------------------------------------------|
    |  0-4294967295  |  Maximum number of packets to allow in excess of rate  |
`,
		},

		"rate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (rate) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum average matching rate

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  integer/unit (Example: 5/minute)  |
`,
			Description: `Maximum average matching rate

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  integer/unit (Example: 5/minute)  |
`,
		},

		// TagNodes

		// Nodes

	}
}
