// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress describes the resource data model.
type ProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddressAsSet       types.Bool   `tfsdk:"as_set" vyos:"as-set,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddressRouteMap    types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddressSummaryOnly types.Bool   `tfsdk:"summary_only" vyos:"summary-only,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress) GetVyosPath() []string {
	return []string{
		"protocols",

		"bgp",

		"address-family",

		"ipv4-multicast",

		"aggregate-address",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `BGP aggregate network/prefix

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  BGP aggregate network/prefix  |

`,
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

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

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

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress) UnmarshalJSON(_ []byte) error {
	return nil
}
