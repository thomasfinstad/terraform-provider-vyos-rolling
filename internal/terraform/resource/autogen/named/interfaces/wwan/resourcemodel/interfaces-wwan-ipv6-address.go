// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWwanIPvsixAddress describes the resource data model.
type InterfacesWwanIPvsixAddress struct {
	// LeafNodes
	LeafInterfacesWwanIPvsixAddressAutoconf           types.Bool `tfsdk:"autoconf" vyos:"autoconf,omitempty"`
	LeafInterfacesWwanIPvsixAddressEuisixfour         types.List `tfsdk:"eui64" vyos:"eui64,omitempty"`
	LeafInterfacesWwanIPvsixAddressNoDefaultLinkLocal types.Bool `tfsdk:"no_default_link_local" vyos:"no-default-link-local,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWwanIPvsixAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"autoconf": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"eui64": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

    |  Format  |  Description  |
    |----------|---------------|
    |  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |

`,
		},

		"no_default_link_local": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Remove the default link-local address from the interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWwanIPvsixAddress) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWwanIPvsixAddress) UnmarshalJSON(_ []byte) error {
	return nil
}
