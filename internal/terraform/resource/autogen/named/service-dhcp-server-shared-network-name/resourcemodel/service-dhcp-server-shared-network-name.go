// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDhcpServerSharedNetworkName describes the resource data model.
type ServiceDhcpServerSharedNetworkName struct {
	// LeafNodes
	ServiceDhcpServerSharedNetworkNameAuthoritative           customtypes.CustomStringValue `tfsdk:"authoritative" json:"authoritative,omitempty"`
	ServiceDhcpServerSharedNetworkNameDomainName              customtypes.CustomStringValue `tfsdk:"domain_name" json:"domain-name,omitempty"`
	ServiceDhcpServerSharedNetworkNameDomainSearch            customtypes.CustomStringValue `tfsdk:"domain_search" json:"domain-search,omitempty"`
	ServiceDhcpServerSharedNetworkNameNtpServer               customtypes.CustomStringValue `tfsdk:"ntp_server" json:"ntp-server,omitempty"`
	ServiceDhcpServerSharedNetworkNamePingCheck               customtypes.CustomStringValue `tfsdk:"ping_check" json:"ping-check,omitempty"`
	ServiceDhcpServerSharedNetworkNameDescrIPtion             customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	ServiceDhcpServerSharedNetworkNameDisable                 customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	ServiceDhcpServerSharedNetworkNameNameServer              customtypes.CustomStringValue `tfsdk:"name_server" json:"name-server,omitempty"`
	ServiceDhcpServerSharedNetworkNameSharedNetworkParameters customtypes.CustomStringValue `tfsdk:"shared_network_parameters" json:"shared-network-parameters,omitempty"`

	// TagNodes
	ServiceDhcpServerSharedNetworkNameSubnet types.Map `tfsdk:"subnet" json:"subnet,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkName) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"authoritative": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Option to make DHCP server authoritative for this physical network

`,
		},

		"domain_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Client Domain Name

`,
		},

		"domain_search": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Client Domain Name search list

`,
		},

		"ntp_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address of NTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  NTP server IPv4 address  |

`,
		},

		"ping_check": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Sends ICMP Echo request to the address being assigned

`,
		},

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable instance

`,
		},

		"name_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |

`,
		},

		"shared_network_parameters": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Additional shared-network parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
		},

		// TagNodes

		"subnet": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnet{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `DHCP subnet for shared network

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |

`,
		},

		// Nodes

	}
}
