/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfMplsTe{}

// VrfNameProtocolsOspfMplsTe describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsOspfMplsTe struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfMplsTeEnable        types.Bool   `tfsdk:"enable" vyos:"enable,omitempty"`
	LeafVrfNameProtocolsOspfMplsTeRouterAddress types.String `tfsdk:"router_address" vyos:"router-address,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfMplsTe) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"enable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable MPLS-TE functionality

`,
			Description: `Enable MPLS-TE functionality

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"router_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Stable IP address of the advertising router

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  ipv4    |  Stable IP address of the advertising router  |
`,
			Description: `Stable IP address of the advertising router

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  ipv4    |  Stable IP address of the advertising router  |
`,

			// Default:          stringdefault.StaticString(`0.0.0.0`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
