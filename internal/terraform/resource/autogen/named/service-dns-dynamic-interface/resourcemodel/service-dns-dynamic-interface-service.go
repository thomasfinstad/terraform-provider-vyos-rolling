// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDNSDynamicInterfaceService describes the resource data model.
type ServiceDNSDynamicInterfaceService struct {
	// LeafNodes
	ServiceDNSDynamicInterfaceServiceHostName customtypes.CustomStringValue `tfsdk:"host_name" json:"host-name,omitempty"`
	ServiceDNSDynamicInterfaceServiceLogin    customtypes.CustomStringValue `tfsdk:"login" json:"login,omitempty"`
	ServiceDNSDynamicInterfaceServicePassword customtypes.CustomStringValue `tfsdk:"password" json:"password,omitempty"`
	ServiceDNSDynamicInterfaceServiceProtocol customtypes.CustomStringValue `tfsdk:"protocol" json:"protocol,omitempty"`
	ServiceDNSDynamicInterfaceServiceServer   customtypes.CustomStringValue `tfsdk:"server" json:"server,omitempty"`
	ServiceDNSDynamicInterfaceServiceZone     customtypes.CustomStringValue `tfsdk:"zone" json:"zone,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDNSDynamicInterfaceService) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"host_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Hostname registered with DDNS service

`,
		},

		"login": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Login for DDNS service

`,
		},

		"password": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Password for DDNS service

`,
		},

		"protocol": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `ddclient protocol used for DDNS service

|  Format  |  Description  |
|----------|---------------|
|  changeip  |  ChangeIP protocol  |
|  cloudflare  |  Cloudflare protocol  |
|  dnsmadeeasy  |  DNS Made Easy protocol  |
|  dnspark  |  DNS Park protocol  |
|  dondominio  |  DonDominio protocol  |
|  dslreports1  |  DslReports protocol  |
|  dtdns  |  DtDNS protocol  |
|  duckdns  |  DuckDNS protocol  |
|  dyndns2  |  DynDNS protocol v2  |
|  easydns  |  easyDNS protocol  |
|  freedns  |  FreeDNS protocol  |
|  freemyip  |  freemyip protocol  |
|  googledomains  |  Google domains protocol  |
|  hammernode1  |  Hammernode protocol  |
|  namecheap  |  Namecheap protocol  |
|  nfsn  |  NearlyFreeSpeech DNS protocol  |
|  noip  |  No-IP protocol  |
|  sitelutions  |  Sitelutions protocol  |
|  woima  |  WOIMA protocol  |
|  yandex  |  Yandex.DNS protocol  |
|  zoneedit1  |  Zoneedit protocol  |

`,
		},

		"server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Remote server to connect to

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Server IPv4 address  |
|  hostname  |  Server hostname/FQDN  |

`,
		},

		"zone": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `DNS zone to update (only available with CloudFlare)

`,
		},

		// TagNodes

		// Nodes

	}
}
