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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (address) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatchIPvsixAddress{}

// PolicyRouteMapRuleMatchIPvsixAddress describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyRouteMapRuleMatchIPvsixAddress struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPvsixAddressAccessList types.String `tfsdk:"access_list" vyos:"access-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixAddressPrefixList types.String `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixAddressPrefixLen  types.Number `tfsdk:"prefix_len" vyos:"prefix-len,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPvsixAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

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

		"prefix_len":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (prefix-len) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 prefix-length to match (can be used for kernel routes only)

    |  Format  |  Description    |
    |----------|-----------------|
    |  0-128   |  Prefix length  |
`,
			Description: `IPv6 prefix-length to match (can be used for kernel routes only)

    |  Format  |  Description    |
    |----------|-----------------|
    |  0-128   |  Prefix length  |
`,
		},

		// TagNodes

		// Nodes

	}
}
