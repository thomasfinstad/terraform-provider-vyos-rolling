// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPaths{}

// VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPaths describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPaths struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPathsEbgp types.Number `tfsdk:"ebgp" vyos:"ebgp,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPathsIbgp types.Number `tfsdk:"ibgp" vyos:"ibgp,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPaths) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ebgp": schema.NumberAttribute{
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

		"ibgp": schema.NumberAttribute{
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

		// Nodes

	}
}
