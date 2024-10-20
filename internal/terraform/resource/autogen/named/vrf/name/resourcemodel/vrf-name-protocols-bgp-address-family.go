/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamily{}

// VrfNameProtocolsBgpAddressFamily describes the resource data model.
type VrfNameProtocolsBgpAddressFamily struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicast        *VrfNameProtocolsBgpAddressFamilyIPvfourUnicast        `tfsdk:"ipv4_unicast" vyos:"ipv4-unicast,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourMulticast      *VrfNameProtocolsBgpAddressFamilyIPvfourMulticast      `tfsdk:"ipv4_multicast" vyos:"ipv4-multicast,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast *VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast `tfsdk:"ipv4_labeled_unicast" vyos:"ipv4-labeled-unicast,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourFlowspec       *VrfNameProtocolsBgpAddressFamilyIPvfourFlowspec       `tfsdk:"ipv4_flowspec" vyos:"ipv4-flowspec,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourVpn            *VrfNameProtocolsBgpAddressFamilyIPvfourVpn            `tfsdk:"ipv4_vpn" vyos:"ipv4-vpn,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicast         *VrfNameProtocolsBgpAddressFamilyIPvsixUnicast         `tfsdk:"ipv6_unicast" vyos:"ipv6-unicast,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixMulticast       *VrfNameProtocolsBgpAddressFamilyIPvsixMulticast       `tfsdk:"ipv6_multicast" vyos:"ipv6-multicast,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast  *VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast  `tfsdk:"ipv6_labeled_unicast" vyos:"ipv6-labeled-unicast,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspec        *VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec        `tfsdk:"ipv6_flowspec" vyos:"ipv6-flowspec,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixVpn             *VrfNameProtocolsBgpAddressFamilyIPvsixVpn             `tfsdk:"ipv6_vpn" vyos:"ipv6-vpn,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpn           *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn           `tfsdk:"l2vpn_evpn" vyos:"l2vpn-evpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamily) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"ipv4_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 BGP settings

`,
			Description: `IPv4 BGP settings

`,
		},

		"ipv4_multicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourMulticast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Multicast IPv4 BGP settings

`,
			Description: `Multicast IPv4 BGP settings

`,
		},

		"ipv4_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Labeled Unicast IPv4 BGP settings

`,
			Description: `Labeled Unicast IPv4 BGP settings

`,
		},

		"ipv4_flowspec": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourFlowspec{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Flowspec IPv4 BGP settings

`,
			Description: `Flowspec IPv4 BGP settings

`,
		},

		"ipv4_vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourVpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Unicast VPN IPv4 BGP settings

`,
			Description: `Unicast VPN IPv4 BGP settings

`,
		},

		"ipv6_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 BGP settings

`,
			Description: `IPv6 BGP settings

`,
		},

		"ipv6_multicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixMulticast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Multicast IPv6 BGP settings

`,
			Description: `Multicast IPv6 BGP settings

`,
		},

		"ipv6_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Labeled Unicast IPv6 BGP settings

`,
			Description: `Labeled Unicast IPv6 BGP settings

`,
		},

		"ipv6_flowspec": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Flowspec IPv6 BGP settings

`,
			Description: `Flowspec IPv6 BGP settings

`,
		},

		"ipv6_vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixVpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Unicast VPN IPv6 BGP settings

`,
			Description: `Unicast VPN IPv6 BGP settings

`,
		},

		"l2vpn_evpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `L2VPN EVPN BGP settings

`,
			Description: `L2VPN EVPN BGP settings

`,
		},
	}
}
