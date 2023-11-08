// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceDNSDynamicInterfaceService describes the resource data model.
type ServiceDNSDynamicInterfaceService struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"service_id" vyos:",self-id"`

	ParentIDServiceDNSDynamicInterface types.String `tfsdk:"interface" vyos:"interface,parent-id"`

	// LeafNodes
	LeafServiceDNSDynamicInterfaceServiceHostName types.List   `tfsdk:"host_name" vyos:"host-name,omitempty"`
	LeafServiceDNSDynamicInterfaceServiceLogin    types.String `tfsdk:"login" vyos:"login,omitempty"`
	LeafServiceDNSDynamicInterfaceServicePassword types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafServiceDNSDynamicInterfaceServiceProtocol types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafServiceDNSDynamicInterfaceServiceServer   types.String `tfsdk:"server" vyos:"server,omitempty"`
	LeafServiceDNSDynamicInterfaceServiceZone     types.String `tfsdk:"zone" vyos:"zone,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o ServiceDNSDynamicInterfaceService) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o ServiceDNSDynamicInterfaceService) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSDynamicInterfaceService) GetVyosPath() []string {
	return []string{
		"service",

		"dns",

		"dynamic",

		"interface",
		o.ParentIDServiceDNSDynamicInterface.ValueString(),

		"service",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSDynamicInterfaceService) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"service_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Service being used for Dynamic DNS

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Dynanmic DNS service with a custom name  |
    |  afraid  &emsp; |  afraid.org Services  |
    |  changeip  &emsp; |  changeip.com Services  |
    |  cloudflare  &emsp; |  cloudflare.com Services  |
    |  dnspark  &emsp; |  dnspark.com Services  |
    |  dslreports  &emsp; |  dslreports.com Services  |
    |  dyndns  &emsp; |  dyndns.com Services  |
    |  easydns  &emsp; |  easydns.com Services  |
    |  namecheap  &emsp; |  namecheap.com Services  |
    |  noip  &emsp; |  noip.com Services  |
    |  sitelutions  &emsp; |  sitelutions.com Services  |
    |  zoneedit  &emsp; |  zoneedit.com Services  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface to send DDNS updates for

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"host_name": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Hostname registered with DDNS service

`,
		},

		"login": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Login for DDNS service

`,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password for DDNS service

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ddclient protocol used for DDNS service

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  changeip  &emsp; |  ChangeIP protocol  |
    |  cloudflare  &emsp; |  Cloudflare protocol  |
    |  dnsmadeeasy  &emsp; |  DNS Made Easy protocol  |
    |  dnspark  &emsp; |  DNS Park protocol  |
    |  dondominio  &emsp; |  DonDominio protocol  |
    |  dslreports1  &emsp; |  DslReports protocol  |
    |  dtdns  &emsp; |  DtDNS protocol  |
    |  duckdns  &emsp; |  DuckDNS protocol  |
    |  dyndns2  &emsp; |  DynDNS protocol v2  |
    |  easydns  &emsp; |  easyDNS protocol  |
    |  freedns  &emsp; |  FreeDNS protocol  |
    |  freemyip  &emsp; |  freemyip protocol  |
    |  googledomains  &emsp; |  Google domains protocol  |
    |  hammernode1  &emsp; |  Hammernode protocol  |
    |  namecheap  &emsp; |  Namecheap protocol  |
    |  nfsn  &emsp; |  NearlyFreeSpeech DNS protocol  |
    |  noip  &emsp; |  No-IP protocol  |
    |  sitelutions  &emsp; |  Sitelutions protocol  |
    |  woima  &emsp; |  WOIMA protocol  |
    |  yandex  &emsp; |  Yandex.DNS protocol  |
    |  zoneedit1  &emsp; |  Zoneedit protocol  |

`,
		},

		"server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remote server to connect to

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Server IPv4 address  |
    |  hostname  &emsp; |  Server hostname/FQDN  |

`,
		},

		"zone": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DNS zone to update (only available with CloudFlare)

`,
		},

		// Nodes

	}
}
