// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfvthreeInterfaceBfd describes the resource data model.
type VrfNameProtocolsOspfvthreeInterfaceBfd struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeInterfaceBfdProfile types.String `tfsdk:"profile" vyos:"profile,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeInterfaceBfd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  BFD profile name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfvthreeInterfaceBfd) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfvthreeInterfaceBfd) UnmarshalJSON(_ []byte) error {
	return nil
}
