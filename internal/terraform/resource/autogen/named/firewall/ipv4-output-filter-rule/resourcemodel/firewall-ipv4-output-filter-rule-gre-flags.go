/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputFilterRuleGreFlags{}

// FirewallIPvfourOutputFilterRuleGreFlags describes the resource data model.
type FirewallIPvfourOutputFilterRuleGreFlags struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeFirewallIPvfourOutputFilterRuleGreFlagsKey      *FirewallIPvfourOutputFilterRuleGreFlagsKey      `tfsdk:"key" vyos:"key,omitempty"`
	NodeFirewallIPvfourOutputFilterRuleGreFlagsChecksum *FirewallIPvfourOutputFilterRuleGreFlagsChecksum `tfsdk:"checksum" vyos:"checksum,omitempty"`
	NodeFirewallIPvfourOutputFilterRuleGreFlagsSequence *FirewallIPvfourOutputFilterRuleGreFlagsSequence `tfsdk:"sequence" vyos:"sequence,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputFilterRuleGreFlags) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"key": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourOutputFilterRuleGreFlagsKey{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Header includes optional key field

`,
			Description: `Header includes optional key field

`,
		},

		"checksum": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourOutputFilterRuleGreFlagsChecksum{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Header includes optional checksum

`,
			Description: `Header includes optional checksum

`,
		},

		"sequence": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourOutputFilterRuleGreFlagsSequence{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Header includes a sequence number field

`,
			Description: `Header includes a sequence number field

`,
		},
	}
}
