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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsStaticRoutesixNextHopBfd{}

// VrfNameProtocolsStaticRoutesixNextHopBfd describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsStaticRoutesixNextHopBfd struct {
	// LeafNodes
	LeafVrfNameProtocolsStaticRoutesixNextHopBfdProfile types.String `tfsdk:"profile" vyos:"profile,omitempty"`

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsStaticRoutesixNextHopBfdMultiHop *VrfNameProtocolsStaticRoutesixNextHopBfdMultiHop `tfsdk:"multi_hop" vyos:"multi-hop,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRoutesixNextHopBfd) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
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
			Attributes: VrfNameProtocolsStaticRoutesixNextHopBfdMultiHop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Use BFD multi hop session

`,
			Description: `Use BFD multi hop session

`,
		},
	}
}
