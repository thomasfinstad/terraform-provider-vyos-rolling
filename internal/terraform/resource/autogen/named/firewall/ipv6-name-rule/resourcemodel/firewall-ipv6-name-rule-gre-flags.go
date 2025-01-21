/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (flags) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixNameRuleGreFlags{}

// FirewallIPvsixNameRuleGreFlags describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixNameRuleGreFlags struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeFirewallIPvsixNameRuleGreFlagsKey *FirewallIPvsixNameRuleGreFlagsKey `tfsdk:"key" vyos:"key,omitempty"`

	NodeFirewallIPvsixNameRuleGreFlagsChecksum *FirewallIPvsixNameRuleGreFlagsChecksum `tfsdk:"checksum" vyos:"checksum,omitempty"`

	NodeFirewallIPvsixNameRuleGreFlagsSequence *FirewallIPvsixNameRuleGreFlagsSequence `tfsdk:"sequence" vyos:"sequence,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleGreFlags) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"key": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleGreFlagsKey{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Header includes optional key field

`,
			Description: `Header includes optional key field

`,
		},

		"checksum": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleGreFlagsChecksum{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Header includes optional checksum

`,
			Description: `Header includes optional checksum

`,
		},

		"sequence": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleGreFlagsSequence{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Header includes a sequence number field

`,
			Description: `Header includes a sequence number field

`,
		},
	}
}
