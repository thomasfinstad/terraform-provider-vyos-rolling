// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyAccessListsixRuleSource describes the resource data model.
type PolicyAccessListsixRuleSource struct {
	// LeafNodes
	LeafPolicyAccessListsixRuleSourceAny        types.Bool   `tfsdk:"any" vyos:"any,omitempty"`
	LeafPolicyAccessListsixRuleSourceExactMatch types.Bool   `tfsdk:"exact_match" vyos:"exact-match,omitempty"`
	LeafPolicyAccessListsixRuleSourceNetwork    types.String `tfsdk:"network" vyos:"network,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessListsixRuleSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"any": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Any IP address to match

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"exact_match": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Exact match of the network prefixes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network/netmask to match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |

`,
		},

		// Nodes

	}
}
