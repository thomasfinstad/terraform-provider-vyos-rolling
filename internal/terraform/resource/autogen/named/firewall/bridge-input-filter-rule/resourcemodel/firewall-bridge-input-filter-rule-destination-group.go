// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeInputFilterRuleDestinationGroup{}

// FirewallBrIDgeInputFilterRuleDestinationGroup describes the resource data model.
type FirewallBrIDgeInputFilterRuleDestinationGroup struct {
	// LeafNodes
	LeafFirewallBrIDgeInputFilterRuleDestinationGroupIPvfourAddressGroup types.String `tfsdk:"ipv4_address_group" vyos:"ipv4-address-group,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleDestinationGroupIPvsixAddressGroup  types.String `tfsdk:"ipv6_address_group" vyos:"ipv6-address-group,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleDestinationGroupMacGroup            types.String `tfsdk:"mac_group" vyos:"mac-group,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleDestinationGroupIPvfourNetworkGroup types.String `tfsdk:"ipv4_network_group" vyos:"ipv4-network-group,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleDestinationGroupIPvsixNetworkGroup  types.String `tfsdk:"ipv6_network_group" vyos:"ipv6-network-group,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleDestinationGroupPortGroup           types.String `tfsdk:"port_group" vyos:"port-group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeInputFilterRuleDestinationGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ipv4_address_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of IPv4 addresses

`,
			Description: `Group of IPv4 addresses

`,
		},

		"ipv6_address_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of IPv6 addresses

`,
			Description: `Group of IPv6 addresses

`,
		},

		"mac_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of MAC addresses

`,
			Description: `Group of MAC addresses

`,
		},

		"ipv4_network_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of IPv4 networks

`,
			Description: `Group of IPv4 networks

`,
		},

		"ipv6_network_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of IPv6 networks

`,
			Description: `Group of IPv6 networks

`,
		},

		"port_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of ports

`,
			Description: `Group of ports

`,
		},

		// Nodes

	}
}
