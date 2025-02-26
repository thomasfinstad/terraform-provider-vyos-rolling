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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (bfd) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsStaticRouteNextHopBfd{}

// VrfNameProtocolsStaticRouteNextHopBfd describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsStaticRouteNextHopBfd struct {
	// LeafNodes
	LeafVrfNameProtocolsStaticRouteNextHopBfdProfile types.String `tfsdk:"profile" vyos:"profile,omitempty"`

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsStaticRouteNextHopBfdMultiHop *VrfNameProtocolsStaticRouteNextHopBfdMultiHop `tfsdk:"multi_hop" vyos:"multi-hop,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRouteNextHopBfd) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (profile) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
			Description: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
		},

		// TagNodes

		// Nodes

		"multi_hop": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsStaticRouteNextHopBfdMultiHop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Configure BFD multi-hop session

`,
			Description: `Configure BFD multi-hop session

`,
		},
	}
}
