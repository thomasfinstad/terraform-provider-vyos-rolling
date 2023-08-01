// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsStaticMulticastInterfaceRoute describes the resource data model.
type ProtocolsStaticMulticastInterfaceRoute struct {
	SelfIdentifier types.String `tfsdk:"interface_route_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsStaticMulticastInterfaceRouteNextHopInterface bool `tfsdk:"next_hop_interface" vyos:"next-hop-interface,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticMulticastInterfaceRoute) GetVyosPath() []string {
	return []string{
		"protocols",

		"static",

		"multicast",

		"interface-route",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticMulticastInterfaceRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"interface_route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Multicast interface based route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Network  |

`,
		},

		// LeafNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsStaticMulticastInterfaceRoute) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsStaticMulticastInterfaceRoute) UnmarshalJSON(_ []byte) error {
	return nil
}
