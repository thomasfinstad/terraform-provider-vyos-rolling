/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit{}

// VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitMedium *VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitMedium `tfsdk:"medium" vyos:"medium,omitempty"`

	NodeVrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitHigh *VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitHigh `tfsdk:"high" vyos:"high,omitempty"`

	NodeVrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical *VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical `tfsdk:"critical" vyos:"critical,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"medium": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitMedium{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Compute for critical, high, and medium priority prefixes

`,
			Description: `Compute for critical, high, and medium priority prefixes

`,
		},

		"high": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitHigh{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Compute for critical, and high priority prefixes

`,
			Description: `Compute for critical, and high priority prefixes

`,
		},

		"critical": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Compute for critical priority prefixes only

`,
			Description: `Compute for critical priority prefixes only

`,
		},
	}
}
