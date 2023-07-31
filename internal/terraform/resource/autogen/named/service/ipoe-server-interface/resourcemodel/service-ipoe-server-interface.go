// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceIPoeServerInterface describes the resource data model.
type ServiceIPoeServerInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafServiceIPoeServerInterfaceMode         types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafServiceIPoeServerInterfaceNetwork      types.String `tfsdk:"network" vyos:"network,omitempty"`
	LeafServiceIPoeServerInterfaceClientSubnet types.String `tfsdk:"client_subnet" vyos:"client-subnet,omitempty"`
	LeafServiceIPoeServerInterfaceVlan         types.List   `tfsdk:"vlan" vyos:"vlan,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeServiceIPoeServerInterfaceExternalDhcp *ServiceIPoeServerInterfaceExternalDhcp `tfsdk:"external_dhcp" vyos:"external-dhcp,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerInterface) GetVyosPath() []string {
	return []string{
		"service",

		"ipoe-server",

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface to listen dhcp or unclassified packets

`,
		},

		// LeafNodes

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client connectivity mode

    |  Format  |  Description  |
    |----------|---------------|
    |  l2  |  Client located on same interface as server  |
    |  l3  |  Client located behind a router  |

`,

			// Default:          stringdefault.StaticString(`l2`),
			Computed: true,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enables clients to share the same network or each client has its own vlan

    |  Format  |  Description  |
    |----------|---------------|
    |  shared  |  Multiple clients share the same network  |
    |  vlan  |  One VLAN per client  |

`,

			// Default:          stringdefault.StaticString(`shared`),
			Computed: true,
		},

		"client_subnet": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client address pool

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  IPv4 address and prefix length  |

`,
		},

		"vlan": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `VLAN monitor for automatic creation of VLAN interfaces

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-4094  |  VLAN for automatic creation  |
    |  start-end  |  VLAN range for automatic creation (e.g. 1-4094)  |

`,
		},

		// Nodes

		"external_dhcp": schema.SingleNestedAttribute{
			Attributes: ServiceIPoeServerInterfaceExternalDhcp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP requests will be forwarded

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceIPoeServerInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceIPoeServerInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
