// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsOspfAreaAreaTypeStub describes the resource data model.
type ProtocolsOspfAreaAreaTypeStub struct {
	// LeafNodes
	LeafProtocolsOspfAreaAreaTypeStubDefaultCost types.Number `tfsdk:"default_cost" vyos:"default-cost,omitempty"`
	LeafProtocolsOspfAreaAreaTypeStubNoSummary   types.Bool   `tfsdk:"no_summary" vyos:"no-summary,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAreaAreaTypeStub) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_cost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Summary-default cost

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-16777215  |  Summary default cost  |

`,
		},

		"no_summary": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not inject inter-area routes into the stub

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsOspfAreaAreaTypeStub) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsOspfAreaAreaTypeStub) UnmarshalJSON(_ []byte) error {
	return nil
}
