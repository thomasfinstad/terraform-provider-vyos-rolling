// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesBrIDgeMemberInterface describes the resource data model.
type InterfacesBrIDgeMemberInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDInterfacesBrIDge types.String `tfsdk:"bridge" vyos:"bridge_identifier,parent-id"`

	// LeafNodes
	LeafInterfacesBrIDgeMemberInterfaceNativeVlan  types.Number `tfsdk:"native_vlan" vyos:"native-vlan,omitempty"`
	LeafInterfacesBrIDgeMemberInterfaceAllowedVlan types.List   `tfsdk:"allowed_vlan" vyos:"allowed-vlan,omitempty"`
	LeafInterfacesBrIDgeMemberInterfaceCost        types.Number `tfsdk:"cost" vyos:"cost,omitempty"`
	LeafInterfacesBrIDgeMemberInterfacePriority    types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafInterfacesBrIDgeMemberInterfaceIsolated    types.Bool   `tfsdk:"isolated" vyos:"isolated,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBrIDgeMemberInterface) GetVyosPath() []string {
	return []string{
		"interfaces",

		"bridge",
		o.ParentIDInterfacesBrIDge.ValueString(),

		"member",

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeMemberInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Member interface name

`,
		},

		"bridge_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bridge Interface

    |  Format  |  Description  |
    |----------|---------------|
    |  brN  |  Bridge interface name  |

`,
		},

		// LeafNodes

		"native_vlan": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specify VLAN id which should natively be present on the link

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-4094  |  Virtual Local Area Network (VLAN) ID  |

`,
		},

		"allowed_vlan": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Specify VLAN id which is allowed in this trunk interface

    |  Format  |  Description  |
    |----------|---------------|
    |  <id>  |  VLAN id allowed to pass this interface  |
    |  <idN>-<idM>  |  VLAN id range allowed on this interface (use '-' as delimiter)  |

`,
		},

		"cost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Bridge port cost

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Path cost value for Spanning Tree Protocol  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Bridge port priority

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-63  |  Bridge port priority  |

`,

			// Default:          stringdefault.StaticString(`32`),
			Computed: true,
		},

		"isolated": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Port is isolated (also known as Private-VLAN)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBrIDgeMemberInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesBrIDgeMemberInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
