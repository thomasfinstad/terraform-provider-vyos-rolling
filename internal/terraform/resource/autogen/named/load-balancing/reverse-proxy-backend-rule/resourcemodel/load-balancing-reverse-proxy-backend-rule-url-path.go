// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &LoadBalancingReverseProxyBackendRuleURLPath{}

// LoadBalancingReverseProxyBackendRuleURLPath describes the resource data model.
type LoadBalancingReverseProxyBackendRuleURLPath struct {
	// LeafNodes
	LeafLoadBalancingReverseProxyBackendRuleURLPathBegin types.List `tfsdk:"begin" vyos:"begin,omitempty"`
	LeafLoadBalancingReverseProxyBackendRuleURLPathEnd   types.List `tfsdk:"end" vyos:"end,omitempty"`
	LeafLoadBalancingReverseProxyBackendRuleURLPathExact types.List `tfsdk:"exact" vyos:"exact,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingReverseProxyBackendRuleURLPath) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"begin": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Begin URL match

    |  Format  |  Description  |
    |----------|---------------|
    |  url     |  Begin URL    |
`,
			Description: `Begin URL match

    |  Format  |  Description  |
    |----------|---------------|
    |  url     |  Begin URL    |
`,
		},

		"end": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `End URL match

    |  Format  |  Description  |
    |----------|---------------|
    |  url     |  End URL      |
`,
			Description: `End URL match

    |  Format  |  Description  |
    |----------|---------------|
    |  url     |  End URL      |
`,
		},

		"exact": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Exactly URL match

    |  Format  |  Description  |
    |----------|---------------|
    |  url     |  Exactly URL  |
`,
			Description: `Exactly URL match

    |  Format  |  Description  |
    |----------|---------------|
    |  url     |  Exactly URL  |
`,
		},

		// Nodes

	}
}