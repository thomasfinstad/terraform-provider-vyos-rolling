// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginate{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginate describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginate struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginateIPvfour types.Bool `tfsdk:"ipv4" vyos:"ipv4,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginateIPvsix  types.Bool `tfsdk:"ipv6" vyos:"ipv6,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginate) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ipv4": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 address family

`,
			Description: `IPv4 address family

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ipv6": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 address family

`,
			Description: `IPv6 address family

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
