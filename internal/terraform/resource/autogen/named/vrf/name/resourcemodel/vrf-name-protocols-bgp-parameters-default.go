// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpParametersDefault{}

// VrfNameProtocolsBgpParametersDefault describes the resource data model.
type VrfNameProtocolsBgpParametersDefault struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersDefaultLocalPref types.Number `tfsdk:"local_pref" vyos:"local-pref,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersDefault) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"local_pref": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Default local preference

    |  Format  |  Description       |
    |----------|--------------------|
    |  u32     |  Local preference  |
`,
			Description: `Default local preference

    |  Format  |  Description       |
    |----------|--------------------|
    |  u32     |  Local preference  |
`,
		},

		// Nodes

	}
}
