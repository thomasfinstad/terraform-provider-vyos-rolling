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

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatchIPNexthop{}

// PolicyRouteMapRuleMatchIPNexthop describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyRouteMapRuleMatchIPNexthop struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPNexthopAddress    types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafPolicyRouteMapRuleMatchIPNexthopAccessList types.Number `tfsdk:"access_list" vyos:"access-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPNexthopPrefixLen  types.Number `tfsdk:"prefix_len" vyos:"prefix-len,omitempty"`
	LeafPolicyRouteMapRuleMatchIPNexthopPrefixList types.String `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPNexthopType       types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPNexthop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address to match

    |  Format  |  Description         |
    |----------|----------------------|
    |  ipv4    |  Nexthop IP address  |
`,
			Description: `IP address to match

    |  Format  |  Description         |
    |----------|----------------------|
    |  ipv4    |  Nexthop IP address  |
`,
		},

		"access_list":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (access-list) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IP access-list to match

    |  Format     |  Description                               |
    |-------------|--------------------------------------------|
    |  1-99       |  IP standard access list                   |
    |  100-199    |  IP extended access list                   |
    |  1300-1999  |  IP standard access list (expanded range)  |
    |  2000-2699  |  IP extended access list (expanded range)  |
`,
			Description: `IP access-list to match

    |  Format     |  Description                               |
    |-------------|--------------------------------------------|
    |  1-99       |  IP standard access list                   |
    |  100-199    |  IP extended access list                   |
    |  1300-1999  |  IP standard access list (expanded range)  |
    |  2000-2699  |  IP extended access list (expanded range)  |
`,
		},

		"prefix_len":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (prefix-len) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IP prefix-length to match

    |  Format  |  Description    |
    |----------|-----------------|
    |  0-32    |  Prefix length  |
`,
			Description: `IP prefix-length to match

    |  Format  |  Description    |
    |----------|-----------------|
    |  0-32    |  Prefix length  |
`,
		},

		"prefix_list":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (prefix-list) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP prefix-list to match

`,
			Description: `IP prefix-list to match

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
