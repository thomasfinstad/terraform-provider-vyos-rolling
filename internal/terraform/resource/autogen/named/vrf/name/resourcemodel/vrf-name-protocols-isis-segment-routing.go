/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisSegmentRouting{}

// VrfNameProtocolsIsisSegmentRouting describes the resource data model.
type VrfNameProtocolsIsisSegmentRouting struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisSegmentRoutingMaximumLabelDepth types.Number `tfsdk:"maximum_label_depth" vyos:"maximum-label-depth,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagVrfNameProtocolsIsisSegmentRoutingPrefix bool `tfsdk:"-" vyos:"prefix,child"`

	// Nodes
	NodeVrfNameProtocolsIsisSegmentRoutingGlobalBlock *VrfNameProtocolsIsisSegmentRoutingGlobalBlock `tfsdk:"global_block" vyos:"global-block,omitempty"`
	NodeVrfNameProtocolsIsisSegmentRoutingLocalBlock  *VrfNameProtocolsIsisSegmentRoutingLocalBlock  `tfsdk:"local_block" vyos:"local-block,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisSegmentRouting) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"maximum_label_depth":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum MPLS labels allowed for this router

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-16    |  MPLS label depth  |
`,
			Description: `Maximum MPLS labels allowed for this router

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-16    |  MPLS label depth  |
`,
		},

		// Nodes

		"global_block": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRoutingGlobalBlock{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Segment Routing Global Block label range

`,
			Description: `Segment Routing Global Block label range

`,
		},

		"local_block": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRoutingLocalBlock{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Segment Routing Local Block label range

`,
			Description: `Segment Routing Local Block label range

`,
		},
	}
}
