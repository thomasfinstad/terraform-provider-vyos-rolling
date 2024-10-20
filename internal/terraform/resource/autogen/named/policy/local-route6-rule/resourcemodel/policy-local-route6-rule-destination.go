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

var _ helpers.VyosResourceDataModel = &PolicyLocalRoutesixRuleDestination{}

// PolicyLocalRoutesixRuleDestination describes the resource data model.
type PolicyLocalRoutesixRuleDestination struct {
	// LeafNodes
	LeafPolicyLocalRoutesixRuleDestinationAddress types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafPolicyLocalRoutesixRuleDestinationPort    types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRoutesixRuleDestination) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		// Nodes

	}
}
