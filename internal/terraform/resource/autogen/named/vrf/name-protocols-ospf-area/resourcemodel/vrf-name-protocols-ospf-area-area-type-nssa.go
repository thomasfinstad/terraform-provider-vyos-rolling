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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfAreaAreaTypeNssa{}

// VrfNameProtocolsOspfAreaAreaTypeNssa describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsOspfAreaAreaTypeNssa struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfAreaAreaTypeNssaDefaultCost types.Number `tfsdk:"default_cost" vyos:"default-cost,omitempty"`
	LeafVrfNameProtocolsOspfAreaAreaTypeNssaNoSummary   types.Bool   `tfsdk:"no_summary" vyos:"no-summary,omitempty"`
	LeafVrfNameProtocolsOspfAreaAreaTypeNssaTranSLAte   types.String `tfsdk:"translate" vyos:"translate,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaAreaTypeNssa) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_cost":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Summary-default cost of an NSSA area

    |  Format      |  Description           |
    |--------------|------------------------|
    |  0-16777215  |  Summary default cost  |
`,
			Description: `Summary-default cost of an NSSA area

    |  Format      |  Description           |
    |--------------|------------------------|
    |  0-16777215  |  Summary default cost  |
`,
		},

		"no_summary":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not inject inter-area routes into stub

`,
			Description: `Do not inject inter-area routes into stub

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"translate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Configure NSSA-ABR

    |  Format     |  Description                 |
    |-------------|------------------------------|
    |  always     |  Always translate LSA types  |
    |  candidate  |  Translate for election      |
    |  never      |  Never translate LSA types   |
`,
			Description: `Configure NSSA-ABR

    |  Format     |  Description                 |
    |-------------|------------------------------|
    |  always     |  Always translate LSA types  |
    |  candidate  |  Translate for election      |
    |  never      |  Never translate LSA types   |
`,

			// Default:          stringdefault.StaticString(`candidate`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
