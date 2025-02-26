/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (monitor) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpBmpTargetMonitor{}

// VrfNameProtocolsBgpBmpTargetMonitor describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpBmpTargetMonitor struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsBgpBmpTargetMonitorIPvfourUnicast *VrfNameProtocolsBgpBmpTargetMonitorIPvfourUnicast `tfsdk:"ipv4_unicast" vyos:"ipv4-unicast,omitempty"`

	NodeVrfNameProtocolsBgpBmpTargetMonitorIPvsixUnicast *VrfNameProtocolsBgpBmpTargetMonitorIPvsixUnicast `tfsdk:"ipv6_unicast" vyos:"ipv6-unicast,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpBmpTargetMonitor) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"ipv4_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpBmpTargetMonitorIPvfourUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Address family IPv4 unicast

`,
			Description: `Address family IPv4 unicast

`,
		},

		"ipv6_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpBmpTargetMonitorIPvsixUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Address family IPv6 unicast

`,
			Description: `Address family IPv6 unicast

`,
		},
	}
}
