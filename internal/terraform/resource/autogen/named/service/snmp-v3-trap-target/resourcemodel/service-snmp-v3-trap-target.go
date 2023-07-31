// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceSnmpVthreeTrapTarget describes the resource data model.
type ServiceSnmpVthreeTrapTarget struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafServiceSnmpVthreeTrapTargetPort     types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafServiceSnmpVthreeTrapTargetProtocol types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafServiceSnmpVthreeTrapTargetType     types.String `tfsdk:"type" vyos:"type,omitempty"`
	LeafServiceSnmpVthreeTrapTargetUser     types.String `tfsdk:"user" vyos:"user,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeServiceSnmpVthreeTrapTargetAuth    *ServiceSnmpVthreeTrapTargetAuth    `tfsdk:"auth" vyos:"auth,omitempty"`
	NodeServiceSnmpVthreeTrapTargetPrivacy *ServiceSnmpVthreeTrapTargetPrivacy `tfsdk:"privacy" vyos:"privacy,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSnmpVthreeTrapTarget) GetVyosPath() []string {
	return []string{
		"service",

		"snmp",

		"v3",

		"trap-target",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeTrapTarget) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Defines SNMP target for inform or traps for IP

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  IP address of trap target  |
    |  ipv6  |  IPv6 address of trap target  |

`,
		},

		// LeafNodes

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`162`),
			Computed: true,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to be used (TCP/UDP)

    |  Format  |  Description  |
    |----------|---------------|
    |  udp  |  Listen protocol UDP  |
    |  tcp  |  Listen protocol TCP  |

`,

			// Default:          stringdefault.StaticString(`udp`),
			Computed: true,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the type of notification between inform and trap

    |  Format  |  Description  |
    |----------|---------------|
    |  inform  |  Use INFORM  |
    |  trap  |  Use TRAP  |

`,

			// Default:          stringdefault.StaticString(`inform`),
			Computed: true,
		},

		"user": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines username for authentication

`,
		},

		// Nodes

		"auth": schema.SingleNestedAttribute{
			Attributes: ServiceSnmpVthreeTrapTargetAuth{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Defines the privacy

`,
		},

		"privacy": schema.SingleNestedAttribute{
			Attributes: ServiceSnmpVthreeTrapTargetPrivacy{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Defines the privacy

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceSnmpVthreeTrapTarget) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceSnmpVthreeTrapTarget) UnmarshalJSON(_ []byte) error {
	return nil
}
