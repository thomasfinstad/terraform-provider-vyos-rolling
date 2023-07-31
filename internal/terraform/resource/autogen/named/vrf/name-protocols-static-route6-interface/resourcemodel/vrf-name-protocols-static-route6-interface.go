// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsStaticRoutesixInterface describes the resource data model.
type VrfNameProtocolsStaticRoutesixInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name_identifier,parent-id"`

	ParentIDVrfNameProtocolsStaticRoutesix types.String `tfsdk:"route6" vyos:"route6_identifier,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsStaticRoutesixInterfaceDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVrfNameProtocolsStaticRoutesixInterfaceDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafVrfNameProtocolsStaticRoutesixInterfaceVrf      types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsStaticRoutesixInterface) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"static",

		"route6",
		o.ParentIDVrfNameProtocolsStaticRoutesix.ValueString(),

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRoutesixInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 gateway interface name

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Gateway interface name  |

`,
		},

		"name_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  VRF instance name  |

`,
		},

		"route6_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv6 route

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv6net  |  IPv6 static route  |

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
func (o *VrfNameProtocolsStaticRoutesixInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsStaticRoutesixInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
