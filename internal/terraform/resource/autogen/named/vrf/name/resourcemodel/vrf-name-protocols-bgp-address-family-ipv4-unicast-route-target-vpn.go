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

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpnBoth   types.String `tfsdk:"both" vyos:"both,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpnImport types.String `tfsdk:"import" vyos:"import,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpnExport types.String `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"both":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Target both import and export

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target both import and export

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		"import":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Target import

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target import

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		"export":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Target export

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target export

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		// Nodes

	}
}
