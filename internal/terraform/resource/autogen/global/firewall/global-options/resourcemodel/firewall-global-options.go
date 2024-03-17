// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &FirewallGlobalOptions{}

// FirewallGlobalOptions describes the resource data model.
type FirewallGlobalOptions struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafFirewallGlobalOptionsAllPing                types.String `tfsdk:"all_ping" vyos:"all-ping,omitempty"`
	LeafFirewallGlobalOptionsBroadcastPing          types.String `tfsdk:"broadcast_ping" vyos:"broadcast-ping,omitempty"`
	LeafFirewallGlobalOptionsIPSrcRoute             types.String `tfsdk:"ip_src_route" vyos:"ip-src-route,omitempty"`
	LeafFirewallGlobalOptionsLogMartians            types.String `tfsdk:"log_martians" vyos:"log-martians,omitempty"`
	LeafFirewallGlobalOptionsReceiveRedirects       types.String `tfsdk:"receive_redirects" vyos:"receive-redirects,omitempty"`
	LeafFirewallGlobalOptionsResolverCache          types.Bool   `tfsdk:"resolver_cache" vyos:"resolver-cache,omitempty"`
	LeafFirewallGlobalOptionsResolverInterval       types.Number `tfsdk:"resolver_interval" vyos:"resolver-interval,omitempty"`
	LeafFirewallGlobalOptionsSendRedirects          types.String `tfsdk:"send_redirects" vyos:"send-redirects,omitempty"`
	LeafFirewallGlobalOptionsSourceValIDation       types.String `tfsdk:"source_validation" vyos:"source-validation,omitempty"`
	LeafFirewallGlobalOptionsSynCookies             types.String `tfsdk:"syn_cookies" vyos:"syn-cookies,omitempty"`
	LeafFirewallGlobalOptionsTwaHazardsProtection   types.String `tfsdk:"twa_hazards_protection" vyos:"twa-hazards-protection,omitempty"`
	LeafFirewallGlobalOptionsIPvsixReceiveRedirects types.String `tfsdk:"ipv6_receive_redirects" vyos:"ipv6-receive-redirects,omitempty"`
	LeafFirewallGlobalOptionsIPvsixSourceValIDation types.String `tfsdk:"ipv6_source_validation" vyos:"ipv6-source-validation,omitempty"`
	LeafFirewallGlobalOptionsIPvsixSrcRoute         types.String `tfsdk:"ipv6_src_route" vyos:"ipv6-src-route,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeFirewallGlobalOptionsStatePolicy bool `tfsdk:"-" vyos:"state-policy,child"`
}

// SetID configures the resource ID
func (o *FirewallGlobalOptions) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGlobalOptions) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"global-options",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallGlobalOptions) GetVyosParentPath() []string {
	return []string{
		"firewall",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *FirewallGlobalOptions) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGlobalOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"all_ping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling of all IPv4 ICMP echo requests

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable processing of all IPv4 ICMP echo requests  |
    |  disable  &emsp; |  Disable processing of all IPv4 ICMP echo requests  |

`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		"broadcast_ping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling broadcast IPv4 ICMP echo and timestamp requests

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable processing of broadcast IPv4 ICMP echo/timestamp requests  |
    |  disable  &emsp; |  Disable processing of broadcast IPv4 ICMP echo/timestamp requests  |

`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"ip_src_route": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling IPv4 packets with source route option

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable processing of IPv4 packets with source route option  |
    |  disable  &emsp; |  Disable processing of IPv4 packets with source route option  |

`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"log_martians": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for logging IPv4 packets with invalid addresses

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable logging of IPv4 packets with invalid addresses  |
    |  disable  &emsp; |  Disable logging of Ipv4 packets with invalid addresses  |

`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		"receive_redirects": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling received IPv4 ICMP redirect messages

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable processing of received IPv4 ICMP redirect messages  |
    |  disable  &emsp; |  Disable processing of received IPv4 ICMP redirect messages  |

`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"resolver_cache": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Retains last successful value if domain resolution fails

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"resolver_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Domain resolver update interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 10-3600  &emsp; |  Interval (seconds)  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"send_redirects": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for sending IPv4 ICMP redirect messages

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable sending IPv4 ICMP redirect messages  |
    |  disable  &emsp; |  Disable sending IPv4 ICMP redirect messages  |

`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		"source_validation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for IPv4 source validation by reversed path, as specified in RFC3704

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  strict  &emsp; |  Enable IPv4 Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose  &emsp; |  Enable IPv4 Loose Reverse Path Forwarding as defined in RFC3704  |
    |  disable  &emsp; |  No IPv4 source validation  |

`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"syn_cookies": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for using TCP SYN cookies with IPv4

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable use of TCP SYN cookies with IPv4  |
    |  disable  &emsp; |  Disable use of TCP SYN cookies with IPv4  |

`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		"twa_hazards_protection": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `RFC1337 TCP TIME-WAIT assasination hazards protection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable RFC1337 TIME-WAIT hazards protection  |
    |  disable  &emsp; |  Disable RFC1337 TIME-WAIT hazards protection  |

`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"ipv6_receive_redirects": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling received ICMPv6 redirect messages

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable processing of received ICMPv6 redirect messages  |
    |  disable  &emsp; |  Disable processing of received ICMPv6 redirect messages  |

`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"ipv6_source_validation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for IPv6 source validation by reversed path, as specified in RFC3704

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  strict  &emsp; |  Enable IPv6 Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose  &emsp; |  Enable IPv6 Loose Reverse Path Forwarding as defined in RFC3704  |
    |  disable  &emsp; |  No IPv6 source validation  |

`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"ipv6_src_route": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling IPv6 packets with routing extension header

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable processing of IPv6 packets with routing header type 2  |
    |  disable  &emsp; |  Disable processing of IPv6 packets with routing header  |

`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},
	}
}
