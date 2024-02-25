// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// FirewallIPvsixOutputFilterRuleSynproxy describes the resource data model.
type FirewallIPvsixOutputFilterRuleSynproxy struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixOutputFilterRuleSynproxyTCP *FirewallIPvsixOutputFilterRuleSynproxyTCP `tfsdk:"tcp" vyos:"tcp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputFilterRuleSynproxy) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixOutputFilterRuleSynproxyTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP synproxy options

`,
		},
	}
}
