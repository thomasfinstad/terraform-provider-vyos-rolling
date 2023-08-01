// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvsixNameRuleTCP describes the resource data model.
type FirewallIPvsixNameRuleTCP struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleTCPMss types.String `tfsdk:"mss" vyos:"mss,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixNameRuleTCPFlags *FirewallIPvsixNameRuleTCPFlags `tfsdk:"flags" vyos:"flags,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum segment size (MSS)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-16384  &emsp; |  Maximum segment size  |
    |  <min>-<max>  &emsp; |  TCP MSS range (use '-' as delimiter)  |

`,
		},

		// Nodes

		"flags": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleTCPFlags{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallIPvsixNameRuleTCP) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallIPvsixNameRuleTCP) UnmarshalJSON(_ []byte) error {
	return nil
}