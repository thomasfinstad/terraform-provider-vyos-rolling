// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesPseudoEthernetIPvsixAddress{}

// InterfacesPseudoEthernetIPvsixAddress describes the resource data model.
type InterfacesPseudoEthernetIPvsixAddress struct {
	// LeafNodes
	LeafInterfacesPseudoEthernetIPvsixAddressAutoconf           types.Bool `tfsdk:"autoconf" vyos:"autoconf,omitempty"`
	LeafInterfacesPseudoEthernetIPvsixAddressEuisixfour         types.List `tfsdk:"eui64" vyos:"eui64,omitempty"`
	LeafInterfacesPseudoEthernetIPvsixAddressNoDefaultLinkLocal types.Bool `tfsdk:"no_default_link_local" vyos:"no-default-link-local,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetIPvsixAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"autoconf": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
			Description: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"eui64": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

    |  Format                |  Description       |
    |------------------------|--------------------|
    |  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |
`,
			Description: `Prefix for IPv6 address with MAC-based EUI-64

    |  Format                |  Description       |
    |------------------------|--------------------|
    |  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |
`,
		},

		"no_default_link_local": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Remove the default link-local address from the interface

`,
			Description: `Remove the default link-local address from the interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
