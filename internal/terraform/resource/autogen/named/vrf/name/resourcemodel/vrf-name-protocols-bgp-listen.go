// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpListen describes the resource data model.
type VrfNameProtocolsBgpListen struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpListenLimit types.Number `tfsdk:"limit" vyos:"limit,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpListenRange bool `tfsdk:"range" vyos:"range,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpListen) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of dynamic neighbors that can be created

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-5000  &emsp; |  BGP neighbor limit  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpListen) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpListen) UnmarshalJSON(_ []byte) error {
	return nil
}
