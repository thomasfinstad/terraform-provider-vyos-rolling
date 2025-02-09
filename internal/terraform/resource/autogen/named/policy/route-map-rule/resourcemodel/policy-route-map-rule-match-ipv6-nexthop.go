/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (nexthop) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatchIPvsixNexthop{}

// PolicyRouteMapRuleMatchIPvsixNexthop describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyRouteMapRuleMatchIPvsixNexthop struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPvsixNexthopAddress    types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixNexthopAccessList types.String `tfsdk:"access_list" vyos:"access-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixNexthopPrefixList types.String `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixNexthopType       types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPvsixNexthop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 address of next-hop

    |  Format  |  Description           |
    |----------|------------------------|
    |  ipv6    |  Nexthop IPv6 address  |
`,
			Description: `IPv6 address of next-hop

    |  Format  |  Description           |
    |----------|------------------------|
    |  ipv6    |  Nexthop IPv6 address  |
`,
		},

		"access_list":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (access-list) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 access-list to match

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  IPV6 access list name  |
`,
			Description: `IPv6 access-list to match

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  IPV6 access list name  |
`,
		},

		"prefix_list":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (prefix-list) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 prefix-list to match

`,
			Description: `IPv6 prefix-list to match

`,
		},

		"type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (type) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match type

    |  Format     |  Description  |
    |-------------|---------------|
    |  blackhole  |  Blackhole    |
`,
			Description: `Match type

    |  Format     |  Description  |
    |-------------|---------------|
    |  blackhole  |  Blackhole    |
`,
		},

		// TagNodes

		// Nodes

	}
}
