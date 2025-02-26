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

var _ helpers.VyosResourceDataModel = &FirewallIPvsixForwardFilterRuleGreFlags{}

// FirewallIPvsixForwardFilterRuleGreFlags describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixForwardFilterRuleGreFlags struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeFirewallIPvsixForwardFilterRuleGreFlagsKey *FirewallIPvsixForwardFilterRuleGreFlagsKey `tfsdk:"key" vyos:"key,omitempty"`

	NodeFirewallIPvsixForwardFilterRuleGreFlagsChecksum *FirewallIPvsixForwardFilterRuleGreFlagsChecksum `tfsdk:"checksum" vyos:"checksum,omitempty"`

	NodeFirewallIPvsixForwardFilterRuleGreFlagsSequence *FirewallIPvsixForwardFilterRuleGreFlagsSequence `tfsdk:"sequence" vyos:"sequence,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleGreFlags) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"key": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleGreFlagsKey{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Header includes optional key field

`,
			Description: `Header includes optional key field

`,
		},

		"checksum": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleGreFlagsChecksum{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Header includes optional checksum

`,
			Description: `Header includes optional checksum

`,
		},

		"sequence": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleGreFlagsSequence{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Header includes a sequence number field

`,
			Description: `Header includes a sequence number field

`,
		},
	}
}
