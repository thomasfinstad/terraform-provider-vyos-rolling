// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesTunnelParametersIP{}

// InterfacesTunnelParametersIP describes the resource data model.
type InterfacesTunnelParametersIP struct {
	// LeafNodes
	LeafInterfacesTunnelParametersIPNoPmtuDiscovery types.Bool   `tfsdk:"no_pmtu_discovery" vyos:"no-pmtu-discovery,omitempty"`
	LeafInterfacesTunnelParametersIPIgnoreDf        types.Bool   `tfsdk:"ignore_df" vyos:"ignore-df,omitempty"`
	LeafInterfacesTunnelParametersIPKey             types.Number `tfsdk:"key" vyos:"key,omitempty"`
	LeafInterfacesTunnelParametersIPTos             types.Number `tfsdk:"tos" vyos:"tos,omitempty"`
	LeafInterfacesTunnelParametersIPTTL             types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesTunnelParametersIP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"no_pmtu_discovery": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable path MTU discovery

`,
			Description: `Disable path MTU discovery

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ignore_df": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore the DF (don't fragment) bit

`,
			Description: `Ignore the DF (don't fragment) bit

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"key": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Tunnel key (only GRE tunnels)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32     |  Tunnel key   |
`,
			Description: `Tunnel key (only GRE tunnels)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32     |  Tunnel key   |
`,
		},

		"tos": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specifies TOS value to use in outgoing packets

    |  Format  |  Description            |
    |----------|-------------------------|
    |  0-99    |  Type of Service (TOS)  |
`,
			Description: `Specifies TOS value to use in outgoing packets

    |  Format  |  Description            |
    |----------|-------------------------|
    |  0-99    |  Type of Service (TOS)  |
`,

			// Default:          stringdefault.StaticString(`inherit`),
			Computed: true,
		},

		"ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specifies TTL value to use in outgoing packets

    |  Format  |  Description                                   |
    |----------|------------------------------------------------|
    |  0       |  Inherit - copy value from original IP header  |
    |  1-255   |  Time to Live                                  |
`,
			Description: `Specifies TTL value to use in outgoing packets

    |  Format  |  Description                                   |
    |----------|------------------------------------------------|
    |  0       |  Inherit - copy value from original IP header  |
    |  1-255   |  Time to Live                                  |
`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// Nodes

	}
}