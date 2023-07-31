// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessSecURItyWpaRadiusServer describes the resource data model.
type InterfacesWirelessSecURItyWpaRadiusServer struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDInterfacesWireless types.String `tfsdk:"wireless" vyos:"wireless_identifier,parent-id"`

	// LeafNodes
	LeafInterfacesWirelessSecURItyWpaRadiusServerDisable    types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesWirelessSecURItyWpaRadiusServerKey        types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafInterfacesWirelessSecURItyWpaRadiusServerPort       types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafInterfacesWirelessSecURItyWpaRadiusServerAccounting types.Bool   `tfsdk:"accounting" vyos:"accounting,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWirelessSecURItyWpaRadiusServer) GetVyosPath() []string {
	return []string{
		"interfaces",

		"wireless",
		o.ParentIDInterfacesWireless.ValueString(),

		"security",

		"wpa",

		"radius",

		"server",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessSecURItyWpaRadiusServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `RADIUS server configuration

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  RADIUS server IPv4 address  |

`,
		},

		"wireless_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Wireless (WiFi/WLAN) Network Interface

    |  Format  |  Description  |
    |----------|---------------|
    |  wlanN  |  Wireless (WiFi/WLAN) interface name  |

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

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shared secret key

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Authentication port

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1812`),
			Computed: true,
		},

		"accounting": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable RADIUS server to receive accounting info

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessSecURItyWpaRadiusServer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWirelessSecURItyWpaRadiusServer) UnmarshalJSON(_ []byte) error {
	return nil
}
