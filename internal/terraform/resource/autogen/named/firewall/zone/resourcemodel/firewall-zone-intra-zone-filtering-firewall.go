// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallZoneIntraZoneFilteringFirewall{}

// FirewallZoneIntraZoneFilteringFirewall describes the resource data model.
type FirewallZoneIntraZoneFilteringFirewall struct {
	// LeafNodes
	LeafFirewallZoneIntraZoneFilteringFirewallIPvsixName types.String `tfsdk:"ipv6_name" vyos:"ipv6-name,omitempty"`
	LeafFirewallZoneIntraZoneFilteringFirewallName       types.String `tfsdk:"name" vyos:"name,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZoneIntraZoneFilteringFirewall) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ipv6_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 firewall ruleset

`,
			Description: `IPv6 firewall ruleset

`,
		},

		"name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 firewall ruleset

`,
			Description: `IPv4 firewall ruleset

`,
		},

		// Nodes

	}
}
