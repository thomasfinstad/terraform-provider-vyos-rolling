// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedAsPath  types.Bool `tfsdk:"as_path" vyos:"as-path,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedMed     types.Bool `tfsdk:"med" vyos:"med,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedNextHop types.Bool `tfsdk:"next_hop" vyos:"next-hop,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as_path": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send AS path unchanged

`,
			Description: `Send AS path unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"med": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send multi-exit discriminator unchanged

`,
			Description: `Send multi-exit discriminator unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"next_hop": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send nexthop unchanged

`,
			Description: `Send nexthop unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
