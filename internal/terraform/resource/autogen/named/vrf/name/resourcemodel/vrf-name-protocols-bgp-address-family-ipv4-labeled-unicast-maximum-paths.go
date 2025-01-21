/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (maximum-paths) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPaths{}

// VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPaths describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPaths struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPathsEbgp types.Number `tfsdk:"ebgp" vyos:"ebgp,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPathsIbgp types.Number `tfsdk:"ibgp" vyos:"ibgp,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPaths) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ebgp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ebgp) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `eBGP maximum paths

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  1-256   |  Number of paths to consider  |
`,
			Description: `eBGP maximum paths

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  1-256   |  Number of paths to consider  |
`,
		},

		"ibgp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ibgp) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `iBGP maximum paths

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  1-256   |  Number of paths to consider  |
`,
			Description: `iBGP maximum paths

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  1-256   |  Number of paths to consider  |
`,
		},

		// TagNodes

		// Nodes

	}
}
