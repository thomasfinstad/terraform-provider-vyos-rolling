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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (fragment) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixInputFilterRuleFragment{}

// FirewallIPvsixInputFilterRuleFragment describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixInputFilterRuleFragment struct {
	// LeafNodes
	LeafFirewallIPvsixInputFilterRuleFragmentMatchFrag    types.Bool `tfsdk:"match_frag" vyos:"match-frag,omitempty"`
	LeafFirewallIPvsixInputFilterRuleFragmentMatchNonFrag types.Bool `tfsdk:"match_non_frag" vyos:"match-non-frag,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixInputFilterRuleFragment) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_frag":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (match-frag) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Second and further fragments of fragmented packets

`,
			Description: `Second and further fragments of fragmented packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"match_non_frag":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (match-non-frag) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Head fragments or unfragmented packets

`,
			Description: `Head fragments or unfragmented packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
