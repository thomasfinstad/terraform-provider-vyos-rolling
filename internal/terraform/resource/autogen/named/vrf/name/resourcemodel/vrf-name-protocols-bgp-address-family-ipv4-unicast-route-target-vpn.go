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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (vpn) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpnBoth   types.String `tfsdk:"both" vyos:"both,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpnImport types.String `tfsdk:"import" vyos:"import,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpnExport types.String `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"both":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (both) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (import) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (export) */
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

		// TagNodes

		// Nodes

	}
}
