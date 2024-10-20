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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpParametersBestpathAsPath{}

// VrfNameProtocolsBgpParametersBestpathAsPath describes the resource data model.
type VrfNameProtocolsBgpParametersBestpathAsPath struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersBestpathAsPathConfed         types.Bool `tfsdk:"confed" vyos:"confed,omitempty"`
	LeafVrfNameProtocolsBgpParametersBestpathAsPathIgnore         types.Bool `tfsdk:"ignore" vyos:"ignore,omitempty"`
	LeafVrfNameProtocolsBgpParametersBestpathAsPathMultIPathRelax types.Bool `tfsdk:"multipath_relax" vyos:"multipath-relax,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersBestpathAsPath) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"confed":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Compare AS-path lengths including confederation sets and sequences

`,
			Description: `Compare AS-path lengths including confederation sets and sequences

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ignore":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore AS-path length in selecting a route

`,
			Description: `Ignore AS-path length in selecting a route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"multipath_relax":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allow load sharing across routes that have different AS paths (but same length)

`,
			Description: `Allow load sharing across routes that have different AS paths (but same length)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
