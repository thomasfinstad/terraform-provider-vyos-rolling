// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// HighAvailabilityVirtualServer describes the resource data model.
type HighAvailabilityVirtualServer struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"virtual_server_id" vyos:"-,self-id"`

	// LeafNodes
	LeafHighAvailabilityVirtualServerAddress            types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafHighAvailabilityVirtualServerAlgorithm          types.String `tfsdk:"algorithm" vyos:"algorithm,omitempty"`
	LeafHighAvailabilityVirtualServerDelayLoop          types.Number `tfsdk:"delay_loop" vyos:"delay-loop,omitempty"`
	LeafHighAvailabilityVirtualServerForwardMethod      types.String `tfsdk:"forward_method" vyos:"forward-method,omitempty"`
	LeafHighAvailabilityVirtualServerFwmark             types.Number `tfsdk:"fwmark" vyos:"fwmark,omitempty"`
	LeafHighAvailabilityVirtualServerPort               types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafHighAvailabilityVirtualServerPersistenceTimeout types.Number `tfsdk:"persistence_timeout" vyos:"persistence-timeout,omitempty"`
	LeafHighAvailabilityVirtualServerProtocol           types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagHighAvailabilityVirtualServerRealServer bool `tfsdk:"-" vyos:"real-server,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *HighAvailabilityVirtualServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVirtualServer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

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
			MarkdownDescription: `Load-balancing virtual server alias

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in virtual_server_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  virtual_server_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address  |
    |  ipv6  &emsp; |  IPv6 address  |

`,
		},

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
