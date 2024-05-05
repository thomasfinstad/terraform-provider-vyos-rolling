// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSetLthreevpnNexthop{}

// PolicyRouteMapRuleSetLthreevpnNexthop describes the resource data model.
type PolicyRouteMapRuleSetLthreevpnNexthop struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyRouteMapRuleSetLthreevpnNexthopEncapsulation *PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation `tfsdk:"encapsulation" vyos:"encapsulation,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetLthreevpnNexthop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"encapsulation": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Encapsulation options (for BGP only)

`,
			Description: `Encapsulation options (for BGP only)

`,
		},
	}
}
