/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (local) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisFastRerouteLfaLocal{}

// VrfNameProtocolsIsisFastRerouteLfaLocal describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsIsisFastRerouteLfaLocal struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing *VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing `tfsdk:"load_sharing" vyos:"load-sharing,omitempty"`

	NodeVrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit *VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit `tfsdk:"priority_limit" vyos:"priority-limit,omitempty"`

	// Ignoring Node `VrfNameProtocolsIsisFastRerouteLfaLocalTiebreaker`.
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastRerouteLfaLocal) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"load_sharing": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Load share prefixes across multiple backups

`,
			Description: `Load share prefixes across multiple backups

`,
		},

		"priority_limit": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Limit backup computation up to the prefix priority

`,
			Description: `Limit backup computation up to the prefix priority

`,
		},
	}
}
