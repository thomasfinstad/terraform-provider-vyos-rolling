// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvfourForwardFilterRuleIPsec describes the resource data model.
type FirewallIPvfourForwardFilterRuleIPsec struct {
	// LeafNodes
	LeafFirewallIPvfourForwardFilterRuleIPsecMatchIPsec types.Bool `tfsdk:"match_ipsec" vyos:"match-ipsec,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleIPsecMatchNone  types.Bool `tfsdk:"match_none" vyos:"match-none,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourForwardFilterRuleIPsec) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_ipsec": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound IPsec packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"match_none": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound non-IPsec packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}