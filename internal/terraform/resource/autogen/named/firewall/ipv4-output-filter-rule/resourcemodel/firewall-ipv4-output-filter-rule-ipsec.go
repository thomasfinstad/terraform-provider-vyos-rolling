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

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputFilterRuleIPsec{}

// FirewallIPvfourOutputFilterRuleIPsec describes the resource data model.
type FirewallIPvfourOutputFilterRuleIPsec struct {
	// LeafNodes
	LeafFirewallIPvfourOutputFilterRuleIPsecMatchIPsecOut types.Bool `tfsdk:"match_ipsec_out" vyos:"match-ipsec-out,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleIPsecMatchNoneOut  types.Bool `tfsdk:"match_none_out" vyos:"match-none-out,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputFilterRuleIPsec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_ipsec_out": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Outbound traffic to be IPsec encapsulated

`,
			Description: `Outbound traffic to be IPsec encapsulated

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"match_none_out": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Outbound traffic that will not be IPsec encapsulated

`,
			Description: `Outbound traffic that will not be IPsec encapsulated

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
