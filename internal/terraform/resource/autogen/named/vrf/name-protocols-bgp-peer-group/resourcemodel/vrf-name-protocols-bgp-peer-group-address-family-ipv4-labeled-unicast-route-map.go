// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMap{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMap describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMap struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMapExport types.String `tfsdk:"export" vyos:"export,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMapImport types.String `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastRouteMap) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter outgoing route updates

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Route-map to filter outgoing route updates

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter incoming route updates

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Route-map to filter incoming route updates

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		// Nodes

	}
}
