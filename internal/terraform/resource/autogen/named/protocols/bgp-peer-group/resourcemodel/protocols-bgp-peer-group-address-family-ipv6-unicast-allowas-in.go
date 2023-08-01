// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasInNumber types.Number `tfsdk:"number" vyos:"number,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of occurrences of AS number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-10  &emsp; |  Number of times AS is allowed in path  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn) UnmarshalJSON(_ []byte) error {
	return nil
}
