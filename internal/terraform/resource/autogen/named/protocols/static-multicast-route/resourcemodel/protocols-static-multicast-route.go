// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsStaticMulticastRoute describes the resource data model.
type ProtocolsStaticMulticastRoute struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

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
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticMulticastRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Configure static unicast route into MRIB for multicast RPF lookup

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  Network  |

`,
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
