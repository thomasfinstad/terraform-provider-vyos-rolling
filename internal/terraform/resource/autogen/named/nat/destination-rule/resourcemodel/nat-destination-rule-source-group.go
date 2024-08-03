// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &NatDestinationRuleSourceGroup{}

// NatDestinationRuleSourceGroup describes the resource data model.
type NatDestinationRuleSourceGroup struct {
	// LeafNodes
	LeafNatDestinationRuleSourceGroupAddressGroup types.String `tfsdk:"address_group" vyos:"address-group,omitempty"`
	LeafNatDestinationRuleSourceGroupDomainGroup  types.String `tfsdk:"domain_group" vyos:"domain-group,omitempty"`
	LeafNatDestinationRuleSourceGroupMacGroup     types.String `tfsdk:"mac_group" vyos:"mac-group,omitempty"`
	LeafNatDestinationRuleSourceGroupNetworkGroup types.String `tfsdk:"network_group" vyos:"network-group,omitempty"`
	LeafNatDestinationRuleSourceGroupPortGroup    types.String `tfsdk:"port_group" vyos:"port-group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleSourceGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of addresses

`,
			Description: `Group of addresses

`,
		},

		"domain_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of domains

`,
			Description: `Group of domains

`,
		},

		"mac_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of MAC addresses

`,
			Description: `Group of MAC addresses

`,
		},

		"network_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of networks

`,
			Description: `Group of networks

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
