// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallGroupIPvsixNetworkGroup describes the resource data model.
type FirewallGroupIPvsixNetworkGroup struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"ipv6_network_group_id" vyos:",self-id"`

	// LeafNodes
	LeafFirewallGroupIPvsixNetworkGroupDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallGroupIPvsixNetworkGroupNetwork     types.List   `tfsdk:"network" vyos:"network,omitempty"`
	LeafFirewallGroupIPvsixNetworkGroupInclude     types.List   `tfsdk:"include" vyos:"include,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o FirewallGroupIPvsixNetworkGroup) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o FirewallGroupIPvsixNetworkGroup) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGroupIPvsixNetworkGroup) GetVyosPath() []string {
	return []string{
		"firewall",

		"group",

		"ipv6-network-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGroupIPvsixNetworkGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"ipv6_network_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall ipv6-network-group

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"network": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Network-group member

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 address to match  |

`,
		},

		"include": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Include another ipv6-network-group

`,
		},

		// Nodes

	}
}
