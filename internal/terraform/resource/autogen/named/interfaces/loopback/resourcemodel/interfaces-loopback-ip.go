// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesLoopbackIP describes the resource data model.
type InterfacesLoopbackIP struct {
	// LeafNodes
	LeafInterfacesLoopbackIPSourceValIDation types.String `tfsdk:"source_validation" vyos:"source-validation,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesLoopbackIP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"source_validation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source validation by reversed path (RFC3704)

    |  Format  |  Description  |
    |----------|---------------|
    |  strict  |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose  |  Enable Loose Reverse Path Forwarding as defined in RFC3704  |
    |  disable  |  No source validation  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesLoopbackIP) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesLoopbackIP) UnmarshalJSON(_ []byte) error {
	return nil
}
