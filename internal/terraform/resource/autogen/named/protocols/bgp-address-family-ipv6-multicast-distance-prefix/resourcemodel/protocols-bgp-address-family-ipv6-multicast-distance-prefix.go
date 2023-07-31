// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefixDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix) GetVyosPath() []string {
	return []string{
		"protocols",

		"bgp",

		"address-family",

		"ipv6-multicast",

		"distance",

		"prefix",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Administrative distance for a specific BGP prefix

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv6net  |  Administrative distance for a specific BGP prefix  |

`,
		},

		// LeafNodes

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Administrative distance for prefix

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Administrative distance for external BGP routes  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix) UnmarshalJSON(_ []byte) error {
	return nil
}
