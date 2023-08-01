// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpParametersDefault describes the resource data model.
type VrfNameProtocolsBgpParametersDefault struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersDefaultLocalPref types.Number `tfsdk:"local_pref" vyos:"local-pref,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersDefault) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"local_pref": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Default local preference

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Local preference  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpParametersDefault) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpParametersDefault) UnmarshalJSON(_ []byte) error {
	return nil
}
