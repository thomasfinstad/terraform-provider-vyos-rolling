// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvfourInputFilterRuleAddAddressToGroupDestinationAddress describes the resource data model.
type FirewallIPvfourInputFilterRuleAddAddressToGroupDestinationAddress struct {
	// LeafNodes
	LeafFirewallIPvfourInputFilterRuleAddAddressToGroupDestinationAddressAddressGroup types.String `tfsdk:"address_group" vyos:"address-group,omitempty"`
	LeafFirewallIPvfourInputFilterRuleAddAddressToGroupDestinationAddressTimeout      types.String `tfsdk:"timeout" vyos:"timeout,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourInputFilterRuleAddAddressToGroupDestinationAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Dynamic address-group

`,
		},

		"timeout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set timeout

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>s  &emsp; |  Timeout value in seconds  |
    |  <number>m  &emsp; |  Timeout value in minutes  |
    |  <number>h  &emsp; |  Timeout value in hours  |
    |  <number>d  &emsp; |  Timeout value in days  |

`,
		},

		// Nodes

	}
}
