/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatchIP{}

// PolicyRouteMapRuleMatchIP describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyRouteMapRuleMatchIP struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodePolicyRouteMapRuleMatchIPAddress *PolicyRouteMapRuleMatchIPAddress `tfsdk:"address" vyos:"address,omitempty"`

	NodePolicyRouteMapRuleMatchIPNexthop *PolicyRouteMapRuleMatchIPNexthop `tfsdk:"nexthop" vyos:"nexthop,omitempty"`

	NodePolicyRouteMapRuleMatchIPRouteSource *PolicyRouteMapRuleMatchIPRouteSource `tfsdk:"route_source" vyos:"route-source,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

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
