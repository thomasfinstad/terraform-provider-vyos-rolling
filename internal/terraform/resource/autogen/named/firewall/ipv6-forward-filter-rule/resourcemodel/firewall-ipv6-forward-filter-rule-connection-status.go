// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvsixForwardFilterRuleConnectionStatus describes the resource data model.
type FirewallIPvsixForwardFilterRuleConnectionStatus struct {
	// LeafNodes
	LeafFirewallIPvsixForwardFilterRuleConnectionStatusNat types.String `tfsdk:"nat" vyos:"nat,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleConnectionStatus) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"nat": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAT connection status

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  destination  &emsp; |  Match connections that are subject to destination NAT  |
    |  source  &emsp; |  Match connections that are subject to source NAT  |

`,
		},

		// Nodes

	}
}
