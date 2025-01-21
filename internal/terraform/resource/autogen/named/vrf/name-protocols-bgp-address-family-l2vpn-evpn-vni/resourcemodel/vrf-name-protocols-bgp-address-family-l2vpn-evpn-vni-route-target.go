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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (route-target) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetBoth   types.List `tfsdk:"both" vyos:"both,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetImport types.List `tfsdk:"import" vyos:"import,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTargetExport types.List `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"both":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (both) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Route Target both import and export

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target both import and export

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		"import":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (import) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Route Target import

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
			Description: `Route Target import

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
		},

		"export":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (export) */
		schema.ListAttribute{
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

		// TagNodes

		// Nodes

	}
}
