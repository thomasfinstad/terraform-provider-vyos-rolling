// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTargetExport types.List `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Route Target export

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target export

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		// Nodes

	}
}
