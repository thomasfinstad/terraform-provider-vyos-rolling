// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDNSDynamicInterface describes the resource data model.
type ServiceDNSDynamicInterface struct {
	// LeafNodes
	ServiceDNSDynamicInterfaceIPvsixEnable customtypes.CustomStringValue `tfsdk:"ipv6_enable" json:"ipv6-enable,omitempty"`

	// TagNodes
	ServiceDNSDynamicInterfaceRfctwoonethreesix types.Map `tfsdk:"rfc2136" json:"rfc2136,omitempty"`
	ServiceDNSDynamicInterfaceService           types.Map `tfsdk:"service" json:"service,omitempty"`

	// Nodes
	ServiceDNSDynamicInterfaceUseWeb types.Object `tfsdk:"use_web" json:"use-web,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDNSDynamicInterface) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ipv6_enable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Allow explicit IPv6 addresses for Dynamic DNS for this interface

`,
		},

		// TagNodes

		"rfc2136": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSDynamicInterfaceRfctwoonethreesix{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `RFC2136 Update name

`,
		},

		"service": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSDynamicInterfaceService{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Service being used for Dynamic DNS

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Dynanmic DNS service with a custom name  |
|  afraid  |  afraid.org Services  |
|  changeip  |  changeip.com Services  |
|  cloudflare  |  cloudflare.com Services  |
|  dnspark  |  dnspark.com Services  |
|  dslreports  |  dslreports.com Services  |
|  dyndns  |  dyndns.com Services  |
|  easydns  |  easydns.com Services  |
|  namecheap  |  namecheap.com Services  |
|  noip  |  noip.com Services  |
|  sitelutions  |  sitelutions.com Services  |
|  zoneedit  |  zoneedit.com Services  |

`,
		},

		// Nodes

		"use_web": schema.SingleNestedAttribute{
			Attributes: ServiceDNSDynamicInterfaceUseWeb{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Web check used for obtaining the external IP address

`,
		},
	}
}
