/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (route-target) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn `tfsdk:"vpn" vyos:"vpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Between current address-family and VPN

`,
			Description: `Between current address-family and VPN

`,
		},
	}
}
