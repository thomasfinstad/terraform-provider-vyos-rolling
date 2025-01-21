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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (conditionally-advertise) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnConditionallyAdvertise{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnConditionallyAdvertise describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnConditionallyAdvertise struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnConditionallyAdvertiseAdvertiseMap types.String `tfsdk:"advertise_map" vyos:"advertise-map,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnConditionallyAdvertiseExistMap     types.String `tfsdk:"exist_map" vyos:"exist-map,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnConditionallyAdvertiseNonExistMap  types.String `tfsdk:"non_exist_map" vyos:"non-exist-map,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnConditionallyAdvertise) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"advertise_map":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (advertise-map) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to conditionally advertise routes

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Route-map to conditionally advertise routes

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		"exist_map":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (exist-map) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Advertise routes only if prefixes in exist-map are installed in BGP table

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		"non_exist_map":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (non-exist-map) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
