// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyAccessListRuleDestination{}

// PolicyAccessListRuleDestination describes the resource data model.
type PolicyAccessListRuleDestination struct {
	// LeafNodes
	LeafPolicyAccessListRuleDestinationAny         types.Bool   `tfsdk:"any" vyos:"any,omitempty"`
	LeafPolicyAccessListRuleDestinationHost        types.String `tfsdk:"host" vyos:"host,omitempty"`
	LeafPolicyAccessListRuleDestinationInverseMask types.String `tfsdk:"inverse_mask" vyos:"inverse-mask,omitempty"`
	LeafPolicyAccessListRuleDestinationNetwork     types.String `tfsdk:"network" vyos:"network,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessListRuleDestination) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"any": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Any IP address to match

`,
			Description: `Any IP address to match

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"host": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Single host IP address to match

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  Host address to match  |
`,
			Description: `Single host IP address to match

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  Host address to match  |
`,
		},

		"inverse_mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network/netmask to match (requires network be defined)

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  Inverse-mask to match  |
`,
			Description: `Network/netmask to match (requires network be defined)

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  Inverse-mask to match  |
`,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network/netmask to match (requires inverse-mask be defined)

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4net  |  Inverse-mask to match  |
`,
			Description: `Network/netmask to match (requires inverse-mask be defined)

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4net  |  Inverse-mask to match  |
`,
		},

		// Nodes

	}
}
