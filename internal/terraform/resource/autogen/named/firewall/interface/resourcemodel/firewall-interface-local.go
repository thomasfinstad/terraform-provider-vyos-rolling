// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallInterfaceLocal describes the resource data model.
type FirewallInterfaceLocal struct {
	// LeafNodes
	LeafFirewallInterfaceLocalName       types.String `tfsdk:"name" vyos:"name,omitempty"`
	LeafFirewallInterfaceLocalIPvsixName types.String `tfsdk:"ipv6_name" vyos:"ipv6-name,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallInterfaceLocal) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local IPv4 firewall ruleset name for interface

`,
		},

		"ipv6_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local IPv6 firewall ruleset name for interface

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallInterfaceLocal) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallInterfaceLocal) UnmarshalJSON(_ []byte) error {
	return nil
}
