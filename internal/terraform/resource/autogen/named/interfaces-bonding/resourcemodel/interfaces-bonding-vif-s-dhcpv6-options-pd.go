// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesBondingVifSDhcpvsixOptionsPd describes the resource data model.
type InterfacesBondingVifSDhcpvsixOptionsPd struct {
	// LeafNodes
	InterfacesBondingVifSDhcpvsixOptionsPdLength customtypes.CustomStringValue `tfsdk:"length" json:"length,omitempty"`

	// TagNodes
	InterfacesBondingVifSDhcpvsixOptionsPdInterface types.Map `tfsdk:"interface" json:"interface,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesBondingVifSDhcpvsixOptionsPd) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"length": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Request IPv6 prefix length from peer

|  Format  |  Description  |
|----------|---------------|
|  u32:32-64  |  Length of delegated prefix  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// TagNodes

		"interface": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesBondingVifSDhcpvsixOptionsPdInterface{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
		},

		// Nodes

	}
}
