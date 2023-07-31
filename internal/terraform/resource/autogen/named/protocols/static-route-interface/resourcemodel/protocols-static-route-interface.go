// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsStaticRouteInterface describes the resource data model.
type ProtocolsStaticRouteInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDProtocolsStaticRoute types.String `tfsdk:"route" vyos:"route_identifier,parent-id"`

	// LeafNodes
	LeafProtocolsStaticRouteInterfaceDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafProtocolsStaticRouteInterfaceDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafProtocolsStaticRouteInterfaceVrf      types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticRouteInterface) GetVyosPath() []string {
	return []string{
		"protocols",

		"static",

		"route",
		o.ParentIDProtocolsStaticRoute.ValueString(),

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticRouteInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Next-hop IPv4 router interface

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Gateway interface name  |

`,
		},

		"route_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv4 route

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  IPv4 static route  |

`,
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for this route

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Distance for this route  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF to leak route

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Name of VRF to leak to  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsStaticRouteInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsStaticRouteInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
