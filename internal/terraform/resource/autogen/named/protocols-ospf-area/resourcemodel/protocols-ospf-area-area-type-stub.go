// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsOspfAreaAreaTypeStub describes the resource data model.
type ProtocolsOspfAreaAreaTypeStub struct {
	// LeafNodes
	ProtocolsOspfAreaAreaTypeStubDefaultCost customtypes.CustomStringValue `tfsdk:"default_cost" json:"default-cost,omitempty"`
	ProtocolsOspfAreaAreaTypeStubNoSummary   customtypes.CustomStringValue `tfsdk:"no_summary" json:"no-summary,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsOspfAreaAreaTypeStub) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_cost": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Summary-default cost

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777215  |  Summary default cost  |

`,
		},

		"no_summary": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Do not inject inter-area routes into the stub

`,
		},

		// TagNodes

		// Nodes

	}
}
