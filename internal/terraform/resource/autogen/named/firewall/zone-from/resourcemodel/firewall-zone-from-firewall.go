// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallZoneFromFirewall describes the resource data model.
type FirewallZoneFromFirewall struct {
	// LeafNodes
	LeafFirewallZoneFromFirewallIPvsixName types.String `tfsdk:"ipv6_name" vyos:"ipv6-name,omitempty"`
	LeafFirewallZoneFromFirewallName       types.String `tfsdk:"name" vyos:"name,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZoneFromFirewall) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ipv6_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 firewall ruleset

`,
		},

		"name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 firewall ruleset

`,
		},

		// Nodes

	}
}
