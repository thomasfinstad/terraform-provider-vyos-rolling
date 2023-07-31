// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpAddressFamilyIPvfourMulticastNetwork describes the resource data model.
type ProtocolsBgpAddressFamilyIPvfourMulticastNetwork struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvfourMulticastNetworkBackdoor types.Bool   `tfsdk:"backdoor" vyos:"backdoor,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvfourMulticastNetworkRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastNetwork) GetVyosPath() []string {
	return []string{
		"protocols",

		"bgp",

		"address-family",

		"ipv4-multicast",

		"network",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvfourMulticastNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Import BGP network/prefix into multicast IPv4 RIB

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  Multicast IPv4 BGP network/prefix  |

`,
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

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastNetwork) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastNetwork) UnmarshalJSON(_ []byte) error {
	return nil
}
