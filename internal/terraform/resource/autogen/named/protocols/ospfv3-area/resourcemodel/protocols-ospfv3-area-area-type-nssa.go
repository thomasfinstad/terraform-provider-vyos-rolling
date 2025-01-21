/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (nssa) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsOspfvthreeAreaAreaTypeNssa{}

// ProtocolsOspfvthreeAreaAreaTypeNssa describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsOspfvthreeAreaAreaTypeNssa struct {
	// LeafNodes
	LeafProtocolsOspfvthreeAreaAreaTypeNssaDefaultInformationOriginate types.Bool `tfsdk:"default_information_originate" vyos:"default-information-originate,omitempty"`
	LeafProtocolsOspfvthreeAreaAreaTypeNssaNoSummary                   types.Bool `tfsdk:"no_summary" vyos:"no-summary,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfvthreeAreaAreaTypeNssa) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_information_originate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-information-originate) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Originate Type 7 default into NSSA area

`,
			Description: `Originate Type 7 default into NSSA area

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_summary":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-summary) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not inject inter-area routes into the stub

`,
			Description: `Do not inject inter-area routes into the stub

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
