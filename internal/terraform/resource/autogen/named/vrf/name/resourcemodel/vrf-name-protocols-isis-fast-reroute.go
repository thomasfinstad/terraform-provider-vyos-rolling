/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (fast-reroute) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisFastReroute{}

// VrfNameProtocolsIsisFastReroute describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsIsisFastReroute struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsIsisFastRerouteLfa *VrfNameProtocolsIsisFastRerouteLfa `tfsdk:"lfa" vyos:"lfa,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastReroute) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"lfa": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfa{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Loop free alternate functionality

`,
			Description: `Loop free alternate functionality

`,
		},
	}
}
