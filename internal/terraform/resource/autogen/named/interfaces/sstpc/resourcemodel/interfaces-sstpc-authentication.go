// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesSstpcAuthentication describes the resource data model.
type InterfacesSstpcAuthentication struct {
	// LeafNodes
	LeafInterfacesSstpcAuthenticationUsername types.String `tfsdk:"username" vyos:"username,omitempty"`
	LeafInterfacesSstpcAuthenticationPassword types.String `tfsdk:"password" vyos:"password,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesSstpcAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"username": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Username  |

`,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Password  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesSstpcAuthentication) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesSstpcAuthentication) UnmarshalJSON(_ []byte) error {
	return nil
}
