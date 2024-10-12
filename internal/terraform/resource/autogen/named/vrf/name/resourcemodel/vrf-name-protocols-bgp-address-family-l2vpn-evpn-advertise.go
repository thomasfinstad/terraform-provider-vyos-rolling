// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfour *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfour `tfsdk:"ipv4" vyos:"ipv4,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsix  *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsix  `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"ipv4": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfour{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 address family

`,
			Description: `IPv4 address family

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 address family

`,
			Description: `IPv6 address family

`,
		},
	}
}
