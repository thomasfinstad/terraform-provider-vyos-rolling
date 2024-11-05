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

var _ helpers.VyosResourceDataModel = &FirewallIPvsixInputFilterRuleIPsec{}

// FirewallIPvsixInputFilterRuleIPsec describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixInputFilterRuleIPsec struct {
	// LeafNodes
	LeafFirewallIPvsixInputFilterRuleIPsecMatchIPsecIn types.Bool `tfsdk:"match_ipsec_in" vyos:"match-ipsec-in,omitempty"`
	LeafFirewallIPvsixInputFilterRuleIPsecMatchNoneIn  types.Bool `tfsdk:"match_none_in" vyos:"match-none-in,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixInputFilterRuleIPsec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_ipsec_in":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound traffic that was IPsec encapsulated

`,
			Description: `Inbound traffic that was IPsec encapsulated

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"match_none_in":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound traffic that was not IPsec encapsulated

`,
			Description: `Inbound traffic that was not IPsec encapsulated

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
