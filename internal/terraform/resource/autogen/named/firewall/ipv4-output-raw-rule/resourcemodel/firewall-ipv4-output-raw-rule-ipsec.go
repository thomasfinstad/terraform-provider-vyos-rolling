/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (ipsec) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputRawRuleIPsec{}

// FirewallIPvfourOutputRawRuleIPsec describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvfourOutputRawRuleIPsec struct {
	// LeafNodes
	LeafFirewallIPvfourOutputRawRuleIPsecMatchIPsecOut types.Bool `tfsdk:"match_ipsec_out" vyos:"match-ipsec-out,omitempty"`
	LeafFirewallIPvfourOutputRawRuleIPsecMatchNoneOut  types.Bool `tfsdk:"match_none_out" vyos:"match-none-out,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputRawRuleIPsec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_ipsec_out":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (match-ipsec-out) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (match-none-out) */
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
