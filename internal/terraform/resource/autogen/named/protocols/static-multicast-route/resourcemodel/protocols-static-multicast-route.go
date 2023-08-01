// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsStaticMulticastRoute describes the resource data model.
type ProtocolsStaticMulticastRoute struct {
	SelfIdentifier types.String `tfsdk:"route_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsStaticMulticastRouteNextHop bool `tfsdk:"next_hop" vyos:"next-hop,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticMulticastRoute) GetVyosPath() []string {
	return []string{
		"protocols",

		"static",

		"multicast",

		"route",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticMulticastRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `route_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Configure static unicast route into MRIB for multicast RPF lookup

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Network  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsStaticMulticastRoute) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsStaticMulticastRoute) UnmarshalJSON(_ []byte) error {
	return nil
}
