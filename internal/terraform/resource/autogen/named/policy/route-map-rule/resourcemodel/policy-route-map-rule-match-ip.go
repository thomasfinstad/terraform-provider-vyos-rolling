// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatchIP{}

// PolicyRouteMapRuleMatchIP describes the resource data model.
type PolicyRouteMapRuleMatchIP struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodePolicyRouteMapRuleMatchIPAddress     *PolicyRouteMapRuleMatchIPAddress     `tfsdk:"address" vyos:"address,omitempty"`
	NodePolicyRouteMapRuleMatchIPNexthop     *PolicyRouteMapRuleMatchIPNexthop     `tfsdk:"nexthop" vyos:"nexthop,omitempty"`
	NodePolicyRouteMapRuleMatchIPRouteSource *PolicyRouteMapRuleMatchIPRouteSource `tfsdk:"route_source" vyos:"route-source,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IP address of route to match

`,
			Description: `IP address of route to match

`,
		},

		"nexthop": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPNexthop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IP next-hop of route to match

`,
			Description: `IP next-hop of route to match

`,
		},

		"route_source": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPRouteSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match advertising source address of route

`,
			Description: `Match advertising source address of route

`,
		},
	}
}
