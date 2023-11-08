// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsStaticRoute describes the resource data model.
type VrfNameProtocolsStaticRoute struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"route_id" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsStaticRouteDhcpInterface types.String `tfsdk:"dhcp_interface" vyos:"dhcp-interface,omitempty"`
	LeafVrfNameProtocolsStaticRouteDescrIPtion   types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsStaticRouteInterface bool `tfsdk:"interface" vyos:"interface,child"`
	ExistsTagVrfNameProtocolsStaticRouteNextHop   bool `tfsdk:"next_hop" vyos:"next-hop,child"`

	// Nodes
	NodeVrfNameProtocolsStaticRouteBlackhole *VrfNameProtocolsStaticRouteBlackhole `tfsdk:"blackhole" vyos:"blackhole,omitempty"`
	NodeVrfNameProtocolsStaticRouteReject    *VrfNameProtocolsStaticRouteReject    `tfsdk:"reject" vyos:"reject,omitempty"`
}

// GetID returns the resource ID
func (o VrfNameProtocolsStaticRoute) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o VrfNameProtocolsStaticRoute) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsStaticRoute) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"static",

		"route",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv4 route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 static route  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"dhcp_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP interface supplying next-hop IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  DHCP interface name  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		// Nodes

		"blackhole": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsStaticRouteBlackhole{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Silently discard pkts when matched

`,
		},

		"reject": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsStaticRouteReject{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Emit an ICMP unreachable when matched

`,
		},
	}
}
