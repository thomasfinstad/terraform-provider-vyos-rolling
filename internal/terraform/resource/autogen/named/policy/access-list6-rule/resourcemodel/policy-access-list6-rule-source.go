/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyAccessListsixRuleSource{}

// PolicyAccessListsixRuleSource describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyAccessListsixRuleSource struct {
	// LeafNodes
	LeafPolicyAccessListsixRuleSourceAny        types.Bool   `tfsdk:"any" vyos:"any,omitempty"`
	LeafPolicyAccessListsixRuleSourceExactMatch types.Bool   `tfsdk:"exact_match" vyos:"exact-match,omitempty"`
	LeafPolicyAccessListsixRuleSourceNetwork    types.String `tfsdk:"network" vyos:"network,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessListsixRuleSource) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"any":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Any IP address to match

`,
			Description: `Any IP address to match

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"exact_match":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Exact match of the network prefixes

`,
			Description: `Exact match of the network prefixes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"network":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network/netmask to match

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `Network/netmask to match

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
		},

		// TagNodes

		// Nodes

	}
}
