// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfvthreeAreaAreaType{}

// VrfNameProtocolsOspfvthreeAreaAreaType describes the resource data model.
type VrfNameProtocolsOspfvthreeAreaAreaType struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfvthreeAreaAreaTypeNssa *VrfNameProtocolsOspfvthreeAreaAreaTypeNssa `tfsdk:"nssa" vyos:"nssa,omitempty"`
	NodeVrfNameProtocolsOspfvthreeAreaAreaTypeStub *VrfNameProtocolsOspfvthreeAreaAreaTypeStub `tfsdk:"stub" vyos:"stub,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeAreaAreaType) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"nssa": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeAreaAreaTypeNssa{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `NSSA OSPFv3 area

`,
			Description: `NSSA OSPFv3 area

`,
		},

		"stub": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeAreaAreaTypeStub{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Stub OSPFv3 area

`,
			Description: `Stub OSPFv3 area

`,
		},
	}
}
