// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &LoadBalancingReverseProxyServiceRuleSet{}

// LoadBalancingReverseProxyServiceRuleSet describes the resource data model.
type LoadBalancingReverseProxyServiceRuleSet struct {
	// LeafNodes
	LeafLoadBalancingReverseProxyServiceRuleSetRedirectLocation types.String `tfsdk:"redirect_location" vyos:"redirect-location,omitempty"`
	LeafLoadBalancingReverseProxyServiceRuleSetBackend          types.String `tfsdk:"backend" vyos:"backend,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingReverseProxyServiceRuleSet) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"redirect_location": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set URL location

    |  Format  |  Description       |
    |----------|--------------------|
    |  url     |  Set URL location  |
`,
			Description: `Set URL location

    |  Format  |  Description       |
    |----------|--------------------|
    |  url     |  Set URL location  |
`,
		},

		"backend": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Backend name

`,
			Description: `Backend name

`,
		},

		// Nodes

	}
}