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

var _ helpers.VyosResourceDataModel = &FirewallIPvfourNameRuleIPsec{}

// FirewallIPvfourNameRuleIPsec describes the resource data model.
type FirewallIPvfourNameRuleIPsec struct {
	// LeafNodes
	LeafFirewallIPvfourNameRuleIPsecMatchIPsec types.Bool `tfsdk:"match_ipsec" vyos:"match-ipsec,omitempty"`
	LeafFirewallIPvfourNameRuleIPsecMatchNone  types.Bool `tfsdk:"match_none" vyos:"match-none,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourNameRuleIPsec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_ipsec": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound IPsec packets

`,
			Description: `Inbound IPsec packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"match_none": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound non-IPsec packets

`,
			Description: `Inbound non-IPsec packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
