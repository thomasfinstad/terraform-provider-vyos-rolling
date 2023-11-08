// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsStaticMulticastInterfaceRouteNextHopInterface describes the resource data model.
type ProtocolsStaticMulticastInterfaceRouteNextHopInterface struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"next_hop_interface_id" vyos:",self-id"`

	ParentIDProtocolsStaticMulticastInterfaceRoute types.String `tfsdk:"interface_route" vyos:"interface-route,parent-id"`

	// LeafNodes
	LeafProtocolsStaticMulticastInterfaceRouteNextHopInterfaceDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o ProtocolsStaticMulticastInterfaceRouteNextHopInterface) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o ProtocolsStaticMulticastInterfaceRouteNextHopInterface) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticMulticastInterfaceRouteNextHopInterface) GetVyosPath() []string {
	return []string{
		"protocols",

		"static",

		"multicast",

		"interface-route",
		o.ParentIDProtocolsStaticMulticastInterfaceRoute.ValueString(),

		"next-hop-interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticMulticastInterfaceRouteNextHopInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"next_hop_interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Next-hop interface

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"interface_route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Multicast interface based route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Network  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance value for this route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Distance for this route  |

`,
		},

		// Nodes

	}
}
