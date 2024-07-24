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

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSetCommunity{}

// PolicyRouteMapRuleSetCommunity describes the resource data model.
type PolicyRouteMapRuleSetCommunity struct {
	// LeafNodes
	LeafPolicyRouteMapRuleSetCommunityAdd     types.List   `tfsdk:"add" vyos:"add,omitempty"`
	LeafPolicyRouteMapRuleSetCommunityReplace types.List   `tfsdk:"replace" vyos:"replace,omitempty"`
	LeafPolicyRouteMapRuleSetCommunityNone    types.Bool   `tfsdk:"none" vyos:"none,omitempty"`
	LeafPolicyRouteMapRuleSetCommunityDelete  types.String `tfsdk:"delete" vyos:"delete,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetCommunity) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"add": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Add communities to a prefix

    |  Format                      |  Description                                                         |
    |------------------------------|----------------------------------------------------------------------|
    |  <AS:VAL>                    |  Community number in <0-65535:0-65535> format                        |
    |  local-as                    |  Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03         |
    |  no-advertise                |  Well-known communities value NO_ADVERTISE 0xFFFFFF02                |
    |  no-export                   |  Well-known communities value NO_EXPORT 0xFFFFFF01                   |
    |  internet                    |  Well-known communities value 0                                      |
    |  graceful-shutdown           |  Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000           |
    |  accept-own                  |  Well-known communities value ACCEPT_OWN 0xFFFF0001                  |
    |  route-filter-translated-v4  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002  |
    |  route-filter-v4             |  Well-known communities value ROUTE_FILTER_v4 0xFFFF0003             |
    |  route-filter-translated-v6  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004  |
    |  route-filter-v6             |  Well-known communities value ROUTE_FILTER_v6 0xFFFF0005             |
    |  llgr-stale                  |  Well-known communities value LLGR_STALE 0xFFFF0006                  |
    |  no-llgr                     |  Well-known communities value NO_LLGR 0xFFFF0007                     |
    |  accept-own-nexthop          |  Well-known communities value accept-own-nexthop 0xFFFF0008          |
    |  blackhole                   |  Well-known communities value BLACKHOLE 0xFFFF029A                   |
    |  no-peer                     |  Well-known communities value NOPEER 0xFFFFFF04                      |
`,
			Description: `Add communities to a prefix

    |  Format                      |  Description                                                         |
    |------------------------------|----------------------------------------------------------------------|
    |  <AS:VAL>                    |  Community number in <0-65535:0-65535> format                        |
    |  local-as                    |  Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03         |
    |  no-advertise                |  Well-known communities value NO_ADVERTISE 0xFFFFFF02                |
    |  no-export                   |  Well-known communities value NO_EXPORT 0xFFFFFF01                   |
    |  internet                    |  Well-known communities value 0                                      |
    |  graceful-shutdown           |  Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000           |
    |  accept-own                  |  Well-known communities value ACCEPT_OWN 0xFFFF0001                  |
    |  route-filter-translated-v4  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002  |
    |  route-filter-v4             |  Well-known communities value ROUTE_FILTER_v4 0xFFFF0003             |
    |  route-filter-translated-v6  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004  |
    |  route-filter-v6             |  Well-known communities value ROUTE_FILTER_v6 0xFFFF0005             |
    |  llgr-stale                  |  Well-known communities value LLGR_STALE 0xFFFF0006                  |
    |  no-llgr                     |  Well-known communities value NO_LLGR 0xFFFF0007                     |
    |  accept-own-nexthop          |  Well-known communities value accept-own-nexthop 0xFFFF0008          |
    |  blackhole                   |  Well-known communities value BLACKHOLE 0xFFFF029A                   |
    |  no-peer                     |  Well-known communities value NOPEER 0xFFFFFF04                      |
`,
		},

		"replace": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Set communities for a prefix

    |  Format                      |  Description                                                         |
    |------------------------------|----------------------------------------------------------------------|
    |  <AS:VAL>                    |  Community number in <0-65535:0-65535> format                        |
    |  local-as                    |  Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03         |
    |  no-advertise                |  Well-known communities value NO_ADVERTISE 0xFFFFFF02                |
    |  no-export                   |  Well-known communities value NO_EXPORT 0xFFFFFF01                   |
    |  internet                    |  Well-known communities value 0                                      |
    |  graceful-shutdown           |  Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000           |
    |  accept-own                  |  Well-known communities value ACCEPT_OWN 0xFFFF0001                  |
    |  route-filter-translated-v4  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002  |
    |  route-filter-v4             |  Well-known communities value ROUTE_FILTER_v4 0xFFFF0003             |
    |  route-filter-translated-v6  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004  |
    |  route-filter-v6             |  Well-known communities value ROUTE_FILTER_v6 0xFFFF0005             |
    |  llgr-stale                  |  Well-known communities value LLGR_STALE 0xFFFF0006                  |
    |  no-llgr                     |  Well-known communities value NO_LLGR 0xFFFF0007                     |
    |  accept-own-nexthop          |  Well-known communities value accept-own-nexthop 0xFFFF0008          |
    |  blackhole                   |  Well-known communities value BLACKHOLE 0xFFFF029A                   |
    |  no-peer                     |  Well-known communities value NOPEER 0xFFFFFF04                      |
`,
			Description: `Set communities for a prefix

    |  Format                      |  Description                                                         |
    |------------------------------|----------------------------------------------------------------------|
    |  <AS:VAL>                    |  Community number in <0-65535:0-65535> format                        |
    |  local-as                    |  Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03         |
    |  no-advertise                |  Well-known communities value NO_ADVERTISE 0xFFFFFF02                |
    |  no-export                   |  Well-known communities value NO_EXPORT 0xFFFFFF01                   |
    |  internet                    |  Well-known communities value 0                                      |
    |  graceful-shutdown           |  Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000           |
    |  accept-own                  |  Well-known communities value ACCEPT_OWN 0xFFFF0001                  |
    |  route-filter-translated-v4  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002  |
    |  route-filter-v4             |  Well-known communities value ROUTE_FILTER_v4 0xFFFF0003             |
    |  route-filter-translated-v6  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004  |
    |  route-filter-v6             |  Well-known communities value ROUTE_FILTER_v6 0xFFFF0005             |
    |  llgr-stale                  |  Well-known communities value LLGR_STALE 0xFFFF0006                  |
    |  no-llgr                     |  Well-known communities value NO_LLGR 0xFFFF0007                     |
    |  accept-own-nexthop          |  Well-known communities value accept-own-nexthop 0xFFFF0008          |
    |  blackhole                   |  Well-known communities value BLACKHOLE 0xFFFF029A                   |
    |  no-peer                     |  Well-known communities value NOPEER 0xFFFFFF04                      |
`,
		},

		"none": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Completely remove communities attribute from a prefix

`,
			Description: `Completely remove communities attribute from a prefix

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"delete": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remove communities defined in a list from a prefix

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Community-list  |
`,
			Description: `Remove communities defined in a list from a prefix

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Community-list  |
`,
		},

		// Nodes

	}
}
