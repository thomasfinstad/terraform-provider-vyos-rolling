// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsRIPInterfaceSend describes the resource data model.
type ProtocolsRIPInterfaceSend struct {
	// LeafNodes
	LeafProtocolsRIPInterfaceSendVersion types.String `tfsdk:"version" vyos:"version,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPInterfaceSend) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"version": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Limit RIP protocol version

    |  Format  |  Description  |
    |----------|---------------|
    |  1  |  Allow RIPv1 only  |
    |  2  |  Allow RIPv2 only  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsRIPInterfaceSend) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsRIPInterfaceSend) UnmarshalJSON(_ []byte) error {
	return nil
}
