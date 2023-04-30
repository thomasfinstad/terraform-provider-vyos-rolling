// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesPseudoEthernetVifSVifCIPvsixAddress describes the resource data model.
type InterfacesPseudoEthernetVifSVifCIPvsixAddress struct {
	// LeafNodes
	InterfacesPseudoEthernetVifSVifCIPvsixAddressAutoconf           customtypes.CustomStringValue `tfsdk:"autoconf" json:"autoconf,omitempty"`
	InterfacesPseudoEthernetVifSVifCIPvsixAddressEuisixfour         customtypes.CustomStringValue `tfsdk:"eui64" json:"eui64,omitempty"`
	InterfacesPseudoEthernetVifSVifCIPvsixAddressNoDefaultLinkLocal customtypes.CustomStringValue `tfsdk:"no_default_link_local" json:"no-default-link-local,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesPseudoEthernetVifSVifCIPvsixAddress) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"autoconf": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
		},

		"eui64": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

|  Format  |  Description  |
|----------|---------------|
|  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |

`,
		},

		"no_default_link_local": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Remove the default link-local address from the interface

`,
		},

		// TagNodes

		// Nodes

	}
}
