// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixForwardFilterRuleGreFlags{}

// FirewallIPvsixForwardFilterRuleGreFlags describes the resource data model.
type FirewallIPvsixForwardFilterRuleGreFlags struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeFirewallIPvsixForwardFilterRuleGreFlagsKey      *FirewallIPvsixForwardFilterRuleGreFlagsKey      `tfsdk:"key" vyos:"key,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleGreFlagsChecksum *FirewallIPvsixForwardFilterRuleGreFlagsChecksum `tfsdk:"checksum" vyos:"checksum,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleGreFlagsSequence *FirewallIPvsixForwardFilterRuleGreFlagsSequence `tfsdk:"sequence" vyos:"sequence,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleGreFlags) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

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
