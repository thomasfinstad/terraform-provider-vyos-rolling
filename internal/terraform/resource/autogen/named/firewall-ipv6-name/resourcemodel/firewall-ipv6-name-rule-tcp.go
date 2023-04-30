// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// FirewallIPvsixNameRuleTCP describes the resource data model.
type FirewallIPvsixNameRuleTCP struct {
	// LeafNodes
	FirewallIPvsixNameRuleTCPMss customtypes.CustomStringValue `tfsdk:"mss" json:"mss,omitempty"`

	// TagNodes

	// Nodes
	FirewallIPvsixNameRuleTCPFlags types.Object `tfsdk:"flags" json:"flags,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o FirewallIPvsixNameRuleTCP) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mss": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum segment size (MSS)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16384  |  Maximum segment size  |
|  <min>-<max>  |  TCP MSS range (use '-' as delimiter)  |

`,
		},

		// TagNodes

		// Nodes

		"flags": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleTCPFlags{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},
	}
}
