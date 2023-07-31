// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpParametersGracefulRestart describes the resource data model.
type VrfNameProtocolsBgpParametersGracefulRestart struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersGracefulRestartStalepathTime types.Number `tfsdk:"stalepath_time" vyos:"stalepath-time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersGracefulRestart) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"stalepath_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum time to hold onto restarting neighbors stale paths

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-3600  |  Hold time in seconds  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpParametersGracefulRestart) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpParametersGracefulRestart) UnmarshalJSON(_ []byte) error {
	return nil
}
