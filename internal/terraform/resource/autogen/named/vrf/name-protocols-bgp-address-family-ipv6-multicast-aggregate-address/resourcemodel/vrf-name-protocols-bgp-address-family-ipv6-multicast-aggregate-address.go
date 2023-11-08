// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"aggregate_address_id" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressAsSet       types.Bool   `tfsdk:"as_set" vyos:"as-set,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressRouteMap    types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressSummaryOnly types.Bool   `tfsdk:"summary_only" vyos:"summary-only,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o VrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o VrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"bgp",

		"address-family",

		"ipv6-multicast",

		"aggregate-address",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"aggregate_address_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `BGP aggregate network/prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  BGP aggregate network/prefix  |

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

		"as_set": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Generate AS-set path information for this aggregate address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},

		"summary_only": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Announce the aggregate summary network only

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
