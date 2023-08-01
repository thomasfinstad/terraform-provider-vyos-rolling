// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvsixNameRuleInboundInterface describes the resource data model.
type FirewallIPvsixNameRuleInboundInterface struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceName  types.String `tfsdk:"interface_name" vyos:"interface-name,omitempty"`
	LeafFirewallIPvsixNameRuleInboundInterfaceInterfaceGroup types.String `tfsdk:"interface_group" vyos:"interface-group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleInboundInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface

`,
		},

		"interface_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface-group

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallIPvsixNameRuleInboundInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallIPvsixNameRuleInboundInterface) UnmarshalJSON(_ []byte) error {
	return nil
}