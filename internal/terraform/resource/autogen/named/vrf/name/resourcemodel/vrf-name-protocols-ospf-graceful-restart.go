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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfGracefulRestart{}

// VrfNameProtocolsOspfGracefulRestart describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsOspfGracefulRestart struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfGracefulRestartGracePeriod types.Number `tfsdk:"grace_period" vyos:"grace-period,omitempty"`

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsOspfGracefulRestartHelper *VrfNameProtocolsOspfGracefulRestartHelper `tfsdk:"helper" vyos:"helper,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfGracefulRestart) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"grace_period":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum length of the grace period

    |  Format  |  Description                                    |
    |----------|-------------------------------------------------|
    |  1-1800  |  Maximum length of the grace period in seconds  |
`,
			Description: `Maximum length of the grace period

    |  Format  |  Description                                    |
    |----------|-------------------------------------------------|
    |  1-1800  |  Maximum length of the grace period in seconds  |
`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"helper": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfGracefulRestartHelper{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `OSPF graceful-restart helpers

`,
			Description: `OSPF graceful-restart helpers

`,
		},
	}
}
