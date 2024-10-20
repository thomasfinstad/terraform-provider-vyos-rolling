/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatchIPvsix{}

// PolicyRouteMapRuleMatchIPvsix describes the resource data model.
type PolicyRouteMapRuleMatchIPvsix struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodePolicyRouteMapRuleMatchIPvsixAddress *PolicyRouteMapRuleMatchIPvsixAddress `tfsdk:"address" vyos:"address,omitempty"`
	NodePolicyRouteMapRuleMatchIPvsixNexthop *PolicyRouteMapRuleMatchIPvsixNexthop `tfsdk:"nexthop" vyos:"nexthop,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPvsixAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 address of route to match

`,
			Description: `IPv6 address of route to match

`,
		},

		"nexthop": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPvsixNexthop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 next-hop of route to match

`,
			Description: `IPv6 next-hop of route to match

`,
		},
	}
}
