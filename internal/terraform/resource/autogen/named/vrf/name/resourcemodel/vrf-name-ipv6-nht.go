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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (nht) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameIPvsixNht{}

// VrfNameIPvsixNht describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameIPvsixNht struct {
	// LeafNodes
	LeafVrfNameIPvsixNhtNoResolveViaDefault types.Bool `tfsdk:"no_resolve_via_default" vyos:"no-resolve-via-default,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameIPvsixNht) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"no_resolve_via_default":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-resolve-via-default) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not resolve via default route

`,
			Description: `Do not resolve via default route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
