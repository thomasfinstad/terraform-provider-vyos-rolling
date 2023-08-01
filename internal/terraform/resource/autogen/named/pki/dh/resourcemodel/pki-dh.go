// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PkiDh describes the resource data model.
type PkiDh struct {
	SelfIdentifier types.String `tfsdk:"dh_id" vyos:",self-id"`

	// LeafNodes
	LeafPkiDhParameters types.String `tfsdk:"parameters" vyos:"parameters,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiDh) GetVyosPath() []string {
	return []string{
		"pki",

		"dh",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiDh) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"dh_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Diffie-Hellman parameters

`,
		},

		// LeafNodes

		"parameters": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DH parameters in PEM format

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PkiDh) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PkiDh) UnmarshalJSON(_ []byte) error {
	return nil
}