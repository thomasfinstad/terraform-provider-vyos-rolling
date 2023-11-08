// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"network_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor types.Bool   `tfsdk:"backdoor" vyos:"backdoor,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) GetVyosPath() []string {
	return []string{
		"protocols",

		"bgp",

		"address-family",

		"ipv6-labeled-unicast",

		"network",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"network_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Import BGP network/prefix into labeled unicast IPv6 RIB

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  Labeled Unicast IPv6 BGP network/prefix  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"backdoor": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use BGP network/prefix as a backdoor route

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

		// Nodes

	}
}
