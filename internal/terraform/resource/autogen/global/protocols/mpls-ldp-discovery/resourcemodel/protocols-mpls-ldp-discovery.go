// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsMplsLdpDiscovery describes the resource data model.
type ProtocolsMplsLdpDiscovery struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsMplsLdpDiscoveryHelloIPvfourHoldtime    types.Number `tfsdk:"hello_ipv4_holdtime" vyos:"hello-ipv4-holdtime,omitempty"`
	LeafProtocolsMplsLdpDiscoveryHelloIPvfourInterval    types.Number `tfsdk:"hello_ipv4_interval" vyos:"hello-ipv4-interval,omitempty"`
	LeafProtocolsMplsLdpDiscoveryHelloIPvsixHoldtime     types.Number `tfsdk:"hello_ipv6_holdtime" vyos:"hello-ipv6-holdtime,omitempty"`
	LeafProtocolsMplsLdpDiscoveryHelloIPvsixInterval     types.Number `tfsdk:"hello_ipv6_interval" vyos:"hello-ipv6-interval,omitempty"`
	LeafProtocolsMplsLdpDiscoverySessionIPvfourHoldtime  types.Number `tfsdk:"session_ipv4_holdtime" vyos:"session-ipv4-holdtime,omitempty"`
	LeafProtocolsMplsLdpDiscoverySessionIPvsixHoldtime   types.Number `tfsdk:"session_ipv6_holdtime" vyos:"session-ipv6-holdtime,omitempty"`
	LeafProtocolsMplsLdpDiscoveryTransportIPvfourAddress types.String `tfsdk:"transport_ipv4_address" vyos:"transport-ipv4-address,omitempty"`
	LeafProtocolsMplsLdpDiscoveryTransportIPvsixAddress  types.String `tfsdk:"transport_ipv6_address" vyos:"transport-ipv6-address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsMplsLdpDiscovery) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpDiscovery) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"mpls",

		"ldp",

		"discovery",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpDiscovery) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"hello_ipv4_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello IPv4 hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |

`,
		},

		"hello_ipv4_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello IPv4 interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |

`,
		},

		"hello_ipv6_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello IPv6 hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |

`,
		},

		"hello_ipv6_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello IPv6 interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |

`,
		},

		"session_ipv4_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session IPv4 hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 15-65535  &emsp; |  Time in seconds  |

`,
		},

		"session_ipv6_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session IPv6 hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 15-65535  &emsp; |  Time in seconds  |

`,
		},

		"transport_ipv4_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Transport IPv4 address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 bind as transport  |

`,
		},

		"transport_ipv6_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Transport IPv6 address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6  &emsp; |  IPv6 bind as transport  |

`,
		},
	}
}