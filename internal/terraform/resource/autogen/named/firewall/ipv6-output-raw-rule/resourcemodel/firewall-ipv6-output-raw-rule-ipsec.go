/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixOutputRawRuleIPsec{}

// FirewallIPvsixOutputRawRuleIPsec describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixOutputRawRuleIPsec struct {
	// LeafNodes
	LeafFirewallIPvsixOutputRawRuleIPsecMatchIPsecOut types.Bool `tfsdk:"match_ipsec_out" vyos:"match-ipsec-out,omitempty"`
	LeafFirewallIPvsixOutputRawRuleIPsecMatchNoneOut  types.Bool `tfsdk:"match_none_out" vyos:"match-none-out,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputRawRuleIPsec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_ipsec_out":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Outbound traffic to be IPsec encapsulated

`,
			Description: `Outbound traffic to be IPsec encapsulated

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"match_none_out":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Outbound traffic that will not be IPsec encapsulated

`,
			Description: `Outbound traffic that will not be IPsec encapsulated

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
