/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsOspfAreaAreaType{}

// ProtocolsOspfAreaAreaType describes the resource data model.
type ProtocolsOspfAreaAreaType struct {
	// LeafNodes
	LeafProtocolsOspfAreaAreaTypeNormal types.Bool `tfsdk:"normal" vyos:"normal,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsOspfAreaAreaTypeNssa *ProtocolsOspfAreaAreaTypeNssa `tfsdk:"nssa" vyos:"nssa,omitempty"`
	NodeProtocolsOspfAreaAreaTypeStub *ProtocolsOspfAreaAreaTypeStub `tfsdk:"stub" vyos:"stub,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAreaAreaType) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"normal":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
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
			Attributes: ProtocolsOspfAreaAreaTypeNssa{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Not-So-Stubby OSPF area

`,
			Description: `Not-So-Stubby OSPF area

`,
		},

		"stub": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfAreaAreaTypeStub{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Stub OSPF area

`,
			Description: `Stub OSPF area

`,
		},
	}
}
