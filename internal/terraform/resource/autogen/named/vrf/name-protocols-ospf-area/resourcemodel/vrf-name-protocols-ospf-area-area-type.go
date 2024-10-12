// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfAreaAreaType{}

// VrfNameProtocolsOspfAreaAreaType describes the resource data model.
type VrfNameProtocolsOspfAreaAreaType struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfAreaAreaTypeNormal types.Bool `tfsdk:"normal" vyos:"normal,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsOspfAreaAreaTypeNssa *VrfNameProtocolsOspfAreaAreaTypeNssa `tfsdk:"nssa" vyos:"nssa,omitempty"`
	NodeVrfNameProtocolsOspfAreaAreaTypeStub *VrfNameProtocolsOspfAreaAreaTypeStub `tfsdk:"stub" vyos:"stub,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaAreaType) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"normal": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Normal OSPF area

`,
			Description: `Normal OSPF area

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"nssa": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfAreaAreaTypeNssa{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Not-So-Stubby OSPF area

`,
			Description: `Not-So-Stubby OSPF area

`,
		},

		"stub": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfAreaAreaTypeStub{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Stub OSPF area

`,
			Description: `Stub OSPF area

`,
		},
	}
}
