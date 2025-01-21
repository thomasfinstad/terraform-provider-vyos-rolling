/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (sid) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn `tfsdk:"vpn" vyos:"vpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Between current VRF and VPN

`,
			Description: `Between current VRF and VPN

`,
		},
	}
}
