// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesDummyIP describes the resource data model.
type InterfacesDummyIP struct {
	// LeafNodes
	LeafInterfacesDummyIPSourceValIDation  types.String `tfsdk:"source_validation" vyos:"source-validation,omitempty"`
	LeafInterfacesDummyIPDisableForwarding types.Bool   `tfsdk:"disable_forwarding" vyos:"disable-forwarding,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesDummyIP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"source_validation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source validation by reversed path (RFC3704)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  strict  &emsp; |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose  &emsp; |  Enable Loose Reverse Path Forwarding as defined in RFC3704  |
    |  disable  &emsp; |  No source validation  |

`,
		},

		"disable_forwarding": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesDummyIP) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesDummyIP) UnmarshalJSON(_ []byte) error {
	return nil
}
