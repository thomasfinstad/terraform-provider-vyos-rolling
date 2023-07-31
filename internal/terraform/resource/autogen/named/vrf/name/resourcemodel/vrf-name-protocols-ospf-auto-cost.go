// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfAutoCost describes the resource data model.
type VrfNameProtocolsOspfAutoCost struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfAutoCostReferenceBandwIDth types.Number `tfsdk:"reference_bandwidth" vyos:"reference-bandwidth,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAutoCost) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"reference_bandwidth": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Reference bandwidth method to assign cost

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-4294967  |  Reference bandwidth cost in Mbits/sec  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfAutoCost) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfAutoCost) UnmarshalJSON(_ []byte) error {
	return nil
}
