// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRouteRuleSourceGroup describes the resource data model.
type PolicyRouteRuleSourceGroup struct {
	// LeafNodes
	LeafPolicyRouteRuleSourceGroupAddressGroup types.String `tfsdk:"address_group" vyos:"address-group,omitempty"`
	LeafPolicyRouteRuleSourceGroupDomainGroup  types.String `tfsdk:"domain_group" vyos:"domain-group,omitempty"`
	LeafPolicyRouteRuleSourceGroupMacGroup     types.String `tfsdk:"mac_group" vyos:"mac-group,omitempty"`
	LeafPolicyRouteRuleSourceGroupNetworkGroup types.String `tfsdk:"network_group" vyos:"network-group,omitempty"`
	LeafPolicyRouteRuleSourceGroupPortGroup    types.String `tfsdk:"port_group" vyos:"port-group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleSourceGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of addresses

`,
		},

		"domain_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of domains

`,
		},

		"mac_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of MAC addresses

`,
		},

		"network_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of networks

`,
		},

		"port_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of ports

`,
		},

		// Nodes

	}
}
