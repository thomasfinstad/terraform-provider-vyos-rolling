/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (address-family) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamily{}

// VrfNameProtocolsBgpPeerGroupAddressFamily describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpPeerGroupAddressFamily struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast `tfsdk:"ipv4_unicast" vyos:"ipv4-unicast,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicast *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicast `tfsdk:"ipv4_labeled_unicast" vyos:"ipv4-labeled-unicast,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpn *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpn `tfsdk:"ipv4_vpn" vyos:"ipv4-vpn,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast `tfsdk:"ipv6_unicast" vyos:"ipv6-unicast,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast `tfsdk:"ipv6_labeled_unicast" vyos:"ipv6-labeled-unicast,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpn *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpn `tfsdk:"ipv6_vpn" vyos:"ipv6-vpn,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn *VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn `tfsdk:"l2vpn_evpn" vyos:"l2vpn-evpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamily) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"ipv4_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 BGP neighbor parameters

`,
			Description: `IPv4 BGP neighbor parameters

`,
		},

		"ipv4_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 Labeled Unicast BGP neighbor parameters

`,
			Description: `IPv4 Labeled Unicast BGP neighbor parameters

`,
		},

		"ipv4_vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 VPN BGP neighbor parameters

`,
			Description: `IPv4 VPN BGP neighbor parameters

`,
		},

		"ipv6_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 BGP neighbor parameters

`,
			Description: `IPv6 BGP neighbor parameters

`,
		},

		"ipv6_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 Labeled Unicast BGP neighbor parameters

`,
			Description: `IPv6 Labeled Unicast BGP neighbor parameters

`,
		},

		"ipv6_vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 VPN BGP neighbor parameters

`,
			Description: `IPv6 VPN BGP neighbor parameters

`,
		},

		"l2vpn_evpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `L2VPN EVPN BGP settings

`,
			Description: `L2VPN EVPN BGP settings

`,
		},
	}
}
