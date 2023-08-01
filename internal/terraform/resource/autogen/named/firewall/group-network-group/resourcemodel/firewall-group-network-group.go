// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallGroupNetworkGroup describes the resource data model.
type FirewallGroupNetworkGroup struct {
	SelfIdentifier types.String `tfsdk:"network_group_id" vyos:",self-id"`

	// LeafNodes
	LeafFirewallGroupNetworkGroupDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallGroupNetworkGroupNetwork     types.List   `tfsdk:"network" vyos:"network,omitempty"`
	LeafFirewallGroupNetworkGroupInclude     types.List   `tfsdk:"include" vyos:"include,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGroupNetworkGroup) GetVyosPath() []string {
	return []string{
		"firewall",

		"group",

		"network-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGroupNetworkGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `network_group_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"network_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall network-group

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
    |  ipv4net  &emsp; |  IPv4 Subnet to match  |

`,
		},

		"include": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Include another network-group

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallGroupNetworkGroup) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallGroupNetworkGroup) UnmarshalJSON(_ []byte) error {
	return nil
}
