// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// InterfacesOpenvpnServerMfa describes the resource data model.
type InterfacesOpenvpnServerMfa struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesOpenvpnServerMfaTotp *InterfacesOpenvpnServerMfaTotp `tfsdk:"totp" vyos:"totp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnServerMfa) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"totp": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnServerMfaTotp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time-based one-time passwords

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesOpenvpnServerMfa) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesOpenvpnServerMfa) UnmarshalJSON(_ []byte) error {
	return nil
}