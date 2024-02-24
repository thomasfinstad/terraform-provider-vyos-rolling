// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceUpnp describes the resource data model.
type ServiceUpnp struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafServiceUpnpFriendlyName    types.String `tfsdk:"friendly_name" vyos:"friendly-name,omitempty"`
	LeafServiceUpnpWanInterface    types.String `tfsdk:"wan_interface" vyos:"wan-interface,omitempty"`
	LeafServiceUpnpWanIP           types.List   `tfsdk:"wan_ip" vyos:"wan-ip,omitempty"`
	LeafServiceUpnpNatPmp          types.Bool   `tfsdk:"nat_pmp" vyos:"nat-pmp,omitempty"`
	LeafServiceUpnpSecureMode      types.Bool   `tfsdk:"secure_mode" vyos:"secure-mode,omitempty"`
	LeafServiceUpnpPresentationURL types.String `tfsdk:"presentation_url" vyos:"presentation-url,omitempty"`
	LeafServiceUpnpListen          types.List   `tfsdk:"listen" vyos:"listen,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceUpnpRule bool `tfsdk:"-" vyos:"rule,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceUpnpPcpLifetime bool `tfsdk:"-" vyos:"pcp-lifetime,ignore,omitempty"`
	ExistsNodeServiceUpnpStun        bool `tfsdk:"-" vyos:"stun,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceUpnp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceUpnp) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"upnp",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceUpnp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"friendly_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of this service

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Friendly name  |

`,
		},

		"wan_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `WAN network interface

`,
		},

		"wan_ip": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `WAN network IP

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address  |
    |  ipv6  &emsp; |  IPv6 address  |

`,
		},

		"nat_pmp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable NAT-PMP support

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"secure_mode": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Secure Mode

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"presentation_url": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Presentation Url

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Presentation Url  |

`,
		},

		"listen": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IP addresses for service to listen on

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <interface>  &emsp; |  Monitor interface address  |
    |  ipv4  &emsp; |  IPv4 address to listen for incoming connections  |
    |  ipv4net  &emsp; |  IPv4 prefix to listen for incoming connections  |
    |  ipv6  &emsp; |  IPv6 address to listen for incoming connections  |
    |  ipv6net  &emsp; |  IPv6 prefix to listen for incoming connections  |

`,
		},
	}
}
