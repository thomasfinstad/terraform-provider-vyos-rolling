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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (source) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyLocalRoutesixRuleSource{}

// PolicyLocalRoutesixRuleSource describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyLocalRoutesixRuleSource struct {
	// LeafNodes
	LeafPolicyLocalRoutesixRuleSourceAddress types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafPolicyLocalRoutesixRuleSourcePort    types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRoutesixRuleSource) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IPv6 address or prefix

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  ipv6     |  Address to match against  |
    |  ipv6net  |  Prefix to match against   |
`,
			Description: `IPv6 address or prefix

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  ipv6     |  Address to match against  |
    |  ipv6net  |  Prefix to match against   |
`,
		},

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		// TagNodes

		// Nodes

	}
}
