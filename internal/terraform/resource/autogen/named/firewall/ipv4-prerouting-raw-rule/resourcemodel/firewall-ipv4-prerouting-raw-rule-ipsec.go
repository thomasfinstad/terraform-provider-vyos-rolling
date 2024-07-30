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

var _ helpers.VyosResourceDataModel = &FirewallIPvfourPreroutingRawRuleIPsec{}

// FirewallIPvfourPreroutingRawRuleIPsec describes the resource data model.
type FirewallIPvfourPreroutingRawRuleIPsec struct {
	// LeafNodes
	LeafFirewallIPvfourPreroutingRawRuleIPsecMatchIPsecIn types.Bool `tfsdk:"match_ipsec_in" vyos:"match-ipsec-in,omitempty"`
	LeafFirewallIPvfourPreroutingRawRuleIPsecMatchNoneIn  types.Bool `tfsdk:"match_none_in" vyos:"match-none-in,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourPreroutingRawRuleIPsec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_ipsec_in": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound traffic that was IPsec encapsulated

`,
			Description: `Inbound traffic that was IPsec encapsulated

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"match_none_in": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound traffic that was not IPsec encapsulated

`,
			Description: `Inbound traffic that was not IPsec encapsulated

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
