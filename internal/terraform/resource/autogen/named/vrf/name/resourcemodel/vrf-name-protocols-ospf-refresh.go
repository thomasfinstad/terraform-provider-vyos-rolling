// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfRefresh describes the resource data model.
type VrfNameProtocolsOspfRefresh struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfRefreshTimers types.Number `tfsdk:"timers" vyos:"timers,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfRefresh) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"timers": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Refresh timer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 10-1800  &emsp; |  Timer value in seconds  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfRefresh) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfRefresh) UnmarshalJSON(_ []byte) error {
	return nil
}
