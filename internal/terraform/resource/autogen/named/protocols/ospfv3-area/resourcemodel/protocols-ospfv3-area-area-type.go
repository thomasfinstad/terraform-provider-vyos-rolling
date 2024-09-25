// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsOspfvthreeAreaAreaType{}

// ProtocolsOspfvthreeAreaAreaType describes the resource data model.
type ProtocolsOspfvthreeAreaAreaType struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsOspfvthreeAreaAreaTypeNssa *ProtocolsOspfvthreeAreaAreaTypeNssa `tfsdk:"nssa" vyos:"nssa,omitempty"`
	NodeProtocolsOspfvthreeAreaAreaTypeStub *ProtocolsOspfvthreeAreaAreaTypeStub `tfsdk:"stub" vyos:"stub,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfvthreeAreaAreaType) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"nssa": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfvthreeAreaAreaTypeNssa{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `NSSA OSPFv3 area

`,
			Description: `NSSA OSPFv3 area

`,
		},

		"stub": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfvthreeAreaAreaTypeStub{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Stub OSPFv3 area

`,
			Description: `Stub OSPFv3 area

`,
		},
	}
}