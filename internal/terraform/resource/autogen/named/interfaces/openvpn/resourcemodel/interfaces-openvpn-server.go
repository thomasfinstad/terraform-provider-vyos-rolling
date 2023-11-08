// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesOpenvpnServer describes the resource data model.
type InterfacesOpenvpnServer struct {
	// LeafNodes
	LeafInterfacesOpenvpnServerDomainName                types.String `tfsdk:"domain_name" vyos:"domain-name,omitempty"`
	LeafInterfacesOpenvpnServerMaxConnections            types.Number `tfsdk:"max_connections" vyos:"max-connections,omitempty"`
	LeafInterfacesOpenvpnServerNameServer                types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`
	LeafInterfacesOpenvpnServerRejectUnconfiguredClients types.Bool   `tfsdk:"reject_unconfigured_clients" vyos:"reject-unconfigured-clients,omitempty"`
	LeafInterfacesOpenvpnServerSubnet                    types.List   `tfsdk:"subnet" vyos:"subnet,omitempty"`
	LeafInterfacesOpenvpnServerTopology                  types.String `tfsdk:"topology" vyos:"topology,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesOpenvpnServerClient    bool `tfsdk:"client" vyos:"client,child"`
	ExistsTagInterfacesOpenvpnServerPushRoute bool `tfsdk:"push_route" vyos:"push-route,child"`

	// Nodes
	NodeInterfacesOpenvpnServerClientIPPool     *InterfacesOpenvpnServerClientIPPool     `tfsdk:"client_ip_pool" vyos:"client-ip-pool,omitempty"`
	NodeInterfacesOpenvpnServerClientIPvsixPool *InterfacesOpenvpnServerClientIPvsixPool `tfsdk:"client_ipv6_pool" vyos:"client-ipv6-pool,omitempty"`
	NodeInterfacesOpenvpnServerMfa              *InterfacesOpenvpnServerMfa              `tfsdk:"mfa" vyos:"mfa,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"domain_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DNS suffix to be pushed to all clients

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Domain Name Server suffix  |

`,
		},

		"max_connections": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of maximum client connections

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4096  &emsp; |  Number of concurrent clients  |

`,
		},

		"name_server": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Domain Name Server (DNS) IPv4 address  |
    |  ipv6  &emsp; |  Domain Name Server (DNS) IPv6 address  |

`,
		},

		"reject_unconfigured_clients": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Reject connections from clients that are not explicitly configured

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"subnet": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Server-mode subnet (from which client IPs are allocated)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 network and prefix length  |
    |  ipv6net  &emsp; |  IPv6 network and prefix length  |

`,
		},

		"topology": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Topology for clients

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  net30  &emsp; |  net30 topology  |
    |  point-to-point  &emsp; |  Point-to-point topology  |
    |  subnet  &emsp; |  Subnet topology  |

`,

			// Default:          stringdefault.StaticString(`net30`),
			Computed: true,
		},

		// Nodes

		"client_ip_pool": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnServerClientIPPool{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Pool of client IPv4 addresses

`,
		},

		"client_ipv6_pool": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnServerClientIPvsixPool{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Pool of client IPv6 addresses

`,
		},

		"mfa": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnServerMfa{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `multi-factor authentication

`,
		},
	}
}
