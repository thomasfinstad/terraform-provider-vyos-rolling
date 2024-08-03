// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpPeerGroupAddressFamily{}

// ProtocolsBgpPeerGroupAddressFamily describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamily struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast        *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast        `tfsdk:"ipv4_unicast" vyos:"ipv4-unicast,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicast *ProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicast `tfsdk:"ipv4_labeled_unicast" vyos:"ipv4-labeled-unicast,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourVpn            *ProtocolsBgpPeerGroupAddressFamilyIPvfourVpn            `tfsdk:"ipv4_vpn" vyos:"ipv4-vpn,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast         *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast         `tfsdk:"ipv6_unicast" vyos:"ipv6-unicast,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast  *ProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast  `tfsdk:"ipv6_labeled_unicast" vyos:"ipv6-labeled-unicast,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixVpn             *ProtocolsBgpPeerGroupAddressFamilyIPvsixVpn             `tfsdk:"ipv6_vpn" vyos:"ipv6-vpn,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn           *ProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn           `tfsdk:"l2vpn_evpn" vyos:"l2vpn-evpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamily) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"ipv4_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 BGP neighbor parameters

`,
			Description: `IPv4 BGP neighbor parameters

`,
		},

		"ipv4_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 Labeled Unicast BGP neighbor parameters

`,
			Description: `IPv4 Labeled Unicast BGP neighbor parameters

`,
		},

		"ipv4_vpn": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourVpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 VPN BGP neighbor parameters

`,
			Description: `IPv4 VPN BGP neighbor parameters

`,
		},

		"ipv6_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 BGP neighbor parameters

`,
			Description: `IPv6 BGP neighbor parameters

`,
		},

		"ipv6_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 Labeled Unicast BGP neighbor parameters

`,
			Description: `IPv6 Labeled Unicast BGP neighbor parameters

`,
		},

		"ipv6_vpn": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixVpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 VPN BGP neighbor parameters

`,
			Description: `IPv6 VPN BGP neighbor parameters

`,
		},

		"l2vpn_evpn": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `L2VPN EVPN BGP settings

`,
			Description: `L2VPN EVPN BGP settings

`,
		},
	}
}
