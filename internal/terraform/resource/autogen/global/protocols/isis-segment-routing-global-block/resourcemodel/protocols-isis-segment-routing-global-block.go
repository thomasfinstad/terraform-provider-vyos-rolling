// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsIsisSegmentRoutingGlobalBlock describes the resource data model.
type ProtocolsIsisSegmentRoutingGlobalBlock struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafProtocolsIsisSegmentRoutingGlobalBlockLowLabelValue  types.Number `tfsdk:"low_label_value" vyos:"low-label-value,omitempty"`
	LeafProtocolsIsisSegmentRoutingGlobalBlockHighLabelValue types.Number `tfsdk:"high_label_value" vyos:"high-label-value,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsIsisSegmentRoutingGlobalBlock) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsIsisSegmentRoutingGlobalBlock) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"isis",

		"segment-routing",

		"global-block",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIsisSegmentRoutingGlobalBlock) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"low_label_value": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label lower bound

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 16-1048575  &emsp; |  Label value (recommended minimum value: 300)  |

`,
		},

		"high_label_value": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label upper bound

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 16-1048575  &emsp; |  Label value  |

`,
		},
	}
}
