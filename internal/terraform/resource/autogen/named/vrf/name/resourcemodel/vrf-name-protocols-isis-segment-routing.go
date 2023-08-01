// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsIsisSegmentRouting describes the resource data model.
type VrfNameProtocolsIsisSegmentRouting struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisSegmentRoutingMaximumLabelDepth types.Number `tfsdk:"maximum_label_depth" vyos:"maximum-label-depth,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsIsisSegmentRoutingPrefix bool `tfsdk:"prefix" vyos:"prefix,child"`

	// Nodes
	NodeVrfNameProtocolsIsisSegmentRoutingGlobalBlock *VrfNameProtocolsIsisSegmentRoutingGlobalBlock `tfsdk:"global_block" vyos:"global-block,omitempty"`
	NodeVrfNameProtocolsIsisSegmentRoutingLocalBlock  *VrfNameProtocolsIsisSegmentRoutingLocalBlock  `tfsdk:"local_block" vyos:"local-block,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisSegmentRouting) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"maximum_label_depth": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum MPLS labels allowed for this router

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-16  &emsp; |  MPLS label depth  |

`,
		},

		// Nodes

		"global_block": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRoutingGlobalBlock{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Segment Routing Global Block label range

`,
		},

		"local_block": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRoutingLocalBlock{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Segment Routing Local Block label range

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisSegmentRouting) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisSegmentRouting) UnmarshalJSON(_ []byte) error {
	return nil
}
