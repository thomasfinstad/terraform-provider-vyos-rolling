// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceSnmpListenAddress describes the resource data model.
type ServiceSnmpListenAddress struct {
	SelfIdentifier types.String `tfsdk:"listen_address_id" vyos:",self-id"`

	// LeafNodes
	LeafServiceSnmpListenAddressPort types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSnmpListenAddress) GetVyosPath() []string {
	return []string{
		"service",

		"snmp",

		"listen-address",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpListenAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"listen_address_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IP address to listen for incoming SNMP requests

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to listen for incoming SNMP requests  |
    |  ipv6  &emsp; |  IPv6 address to listen for incoming SNMP requests  |

`,
		},

		// LeafNodes

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`161`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceSnmpListenAddress) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceSnmpListenAddress) UnmarshalJSON(_ []byte) error {
	return nil
}
