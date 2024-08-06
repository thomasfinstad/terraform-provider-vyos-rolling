// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixNameRuleGreFlags{}

// FirewallIPvsixNameRuleGreFlags describes the resource data model.
type FirewallIPvsixNameRuleGreFlags struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixNameRuleGreFlagsKey      *FirewallIPvsixNameRuleGreFlagsKey      `tfsdk:"key" vyos:"key,omitempty"`
	NodeFirewallIPvsixNameRuleGreFlagsChecksum *FirewallIPvsixNameRuleGreFlagsChecksum `tfsdk:"checksum" vyos:"checksum,omitempty"`
	NodeFirewallIPvsixNameRuleGreFlagsSequence *FirewallIPvsixNameRuleGreFlagsSequence `tfsdk:"sequence" vyos:"sequence,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleGreFlags) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

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
