// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HighAvailabilityVirtualServer describes the resource data model.
type HighAvailabilityVirtualServer struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"virtual_server_id" vyos:",self-id"`

	// LeafNodes
	LeafHighAvailabilityVirtualServerAlgorithm          types.String `tfsdk:"algorithm" vyos:"algorithm,omitempty"`
	LeafHighAvailabilityVirtualServerDelayLoop          types.Number `tfsdk:"delay_loop" vyos:"delay-loop,omitempty"`
	LeafHighAvailabilityVirtualServerForwardMethod      types.String `tfsdk:"forward_method" vyos:"forward-method,omitempty"`
	LeafHighAvailabilityVirtualServerFwmark             types.Number `tfsdk:"fwmark" vyos:"fwmark,omitempty"`
	LeafHighAvailabilityVirtualServerPort               types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafHighAvailabilityVirtualServerPersistenceTimeout types.Number `tfsdk:"persistence_timeout" vyos:"persistence-timeout,omitempty"`
	LeafHighAvailabilityVirtualServerProtocol           types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagHighAvailabilityVirtualServerRealServer bool `tfsdk:"real_server" vyos:"real-server,child"`

	// Nodes
}

// GetID returns the resource ID
func (o HighAvailabilityVirtualServer) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o HighAvailabilityVirtualServer) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVirtualServer) GetVyosPath() []string {
	return []string{
		"high-availability",

		"virtual-server",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVirtualServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"virtual_server_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Load-balancing virtual server address

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"algorithm": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Schedule algorithm (default - least-connection)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  round-robin  &emsp; |  Round robin  |
    |  weighted-round-robin  &emsp; |  Weighted round robin  |
    |  least-connection  &emsp; |  Least connection  |
    |  weighted-least-connection  &emsp; |  Weighted least connection  |
    |  source-hashing  &emsp; |  Source hashing  |
    |  destination-hashing  &emsp; |  Destination hashing  |
    |  locality-based-least-connection  &emsp; |  Locality-Based least connection  |

`,

			// Default:          stringdefault.StaticString(`least-connection`),
			Computed: true,
		},

		"delay_loop": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between health-checks (in seconds)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-600  &emsp; |  Interval in seconds  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"forward_method": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Forwarding method

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  direct  &emsp; |  Direct routing  |
    |  nat  &emsp; |  NAT  |
    |  tunnel  &emsp; |  Tunneling  |

`,

			// Default:          stringdefault.StaticString(`nat`),
			Computed: true,
		},

		"fwmark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match fwmark value

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2147483647  &emsp; |  Match firewall mark value  |

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Numeric IP port  |

`,
		},

		"persistence_timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for persistent connections

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-86400  &emsp; |  Timeout for persistent connections  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol for port checks

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  tcp  &emsp; |  TCP  |
    |  udp  &emsp; |  UDP  |

`,

			// Default:          stringdefault.StaticString(`tcp`),
			Computed: true,
		},

		// Nodes

	}
}
