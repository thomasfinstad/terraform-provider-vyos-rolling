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

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSetIPvsixNextHop{}

// PolicyRouteMapRuleSetIPvsixNextHop describes the resource data model.
type PolicyRouteMapRuleSetIPvsixNextHop struct {
	// LeafNodes
	LeafPolicyRouteMapRuleSetIPvsixNextHopGlobal       types.String `tfsdk:"global" vyos:"global,omitempty"`
	LeafPolicyRouteMapRuleSetIPvsixNextHopLocal        types.String `tfsdk:"local" vyos:"local,omitempty"`
	LeafPolicyRouteMapRuleSetIPvsixNextHopPeerAddress  types.Bool   `tfsdk:"peer_address" vyos:"peer-address,omitempty"`
	LeafPolicyRouteMapRuleSetIPvsixNextHopPreferGlobal types.Bool   `tfsdk:"prefer_global" vyos:"prefer-global,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetIPvsixNextHop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"global": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Nexthop IPv6 global address

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv6    |  IPv6 address and prefix length  |
`,
			Description: `Nexthop IPv6 global address

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv6    |  IPv6 address and prefix length  |
`,
		},

		"local": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Nexthop IPv6 local address

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv6    |  IPv6 address and prefix length  |
`,
			Description: `Nexthop IPv6 local address

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv6    |  IPv6 address and prefix length  |
`,
		},

		"peer_address": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use peer address (for BGP only)

`,
			Description: `Use peer address (for BGP only)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"prefer_global": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Prefer global address as the nexthop

`,
			Description: `Prefer global address as the nexthop

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
