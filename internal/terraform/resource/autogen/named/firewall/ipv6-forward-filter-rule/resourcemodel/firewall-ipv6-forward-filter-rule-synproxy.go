// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixForwardFilterRuleSynproxy{}

// FirewallIPvsixForwardFilterRuleSynproxy describes the resource data model.
type FirewallIPvsixForwardFilterRuleSynproxy struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixForwardFilterRuleSynproxyTCP *FirewallIPvsixForwardFilterRuleSynproxyTCP `tfsdk:"tcp" vyos:"tcp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleSynproxy) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleSynproxyTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP synproxy options

`,
			Description: `TCP synproxy options

`,
		},
	}
}
