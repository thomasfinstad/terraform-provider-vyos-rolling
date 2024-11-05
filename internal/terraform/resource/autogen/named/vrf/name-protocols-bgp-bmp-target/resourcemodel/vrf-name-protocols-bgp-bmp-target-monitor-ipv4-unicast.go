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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpBmpTargetMonitorIPvfourUnicast{}

// VrfNameProtocolsBgpBmpTargetMonitorIPvfourUnicast describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpBmpTargetMonitorIPvfourUnicast struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpBmpTargetMonitorIPvfourUnicastPrePolicy  types.Bool `tfsdk:"pre_policy" vyos:"pre-policy,omitempty"`
	LeafVrfNameProtocolsBgpBmpTargetMonitorIPvfourUnicastPostPolicy types.Bool `tfsdk:"post_policy" vyos:"post-policy,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpBmpTargetMonitorIPvfourUnicast) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"pre_policy":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send state before policy and filter processing

`,
			Description: `Send state before policy and filter processing

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"post_policy":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send state with policy and filters applied

`,
			Description: `Send state with policy and filters applied

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
