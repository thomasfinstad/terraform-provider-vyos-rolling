/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceExternal types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceInternal types.Number `tfsdk:"internal" vyos:"internal,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceLocal    types.Number `tfsdk:"local" vyos:"local,omitempty"`

	// TagNodes

	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix bool `tfsdk:"-" vyos:"prefix,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `eBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  eBGP routes administrative distance  |
`,
			Description: `eBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  eBGP routes administrative distance  |
`,
		},

		"internal":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `iBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  iBGP routes administrative distance  |
`,
			Description: `iBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  iBGP routes administrative distance  |
`,
		},

		"local":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Locally originated BGP routes administrative distance

    |  Format  |  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   |  Locally originated BGP routes administrative distance  |
`,
			Description: `Locally originated BGP routes administrative distance

    |  Format  |  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   |  Locally originated BGP routes administrative distance  |
`,
		},

		// TagNodes

		// Nodes

	}
}
