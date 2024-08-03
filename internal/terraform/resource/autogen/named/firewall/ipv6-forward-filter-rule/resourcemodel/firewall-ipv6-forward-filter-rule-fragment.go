// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixForwardFilterRuleFragment{}

// FirewallIPvsixForwardFilterRuleFragment describes the resource data model.
type FirewallIPvsixForwardFilterRuleFragment struct {
	// LeafNodes
	LeafFirewallIPvsixForwardFilterRuleFragmentMatchFrag    types.Bool `tfsdk:"match_frag" vyos:"match-frag,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleFragmentMatchNonFrag types.Bool `tfsdk:"match_non_frag" vyos:"match-non-frag,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleFragment) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_frag": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Second and further fragments of fragmented packets

`,
			Description: `Second and further fragments of fragmented packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"match_non_frag": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Head fragments or unfragmented packets

`,
			Description: `Head fragments or unfragmented packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
