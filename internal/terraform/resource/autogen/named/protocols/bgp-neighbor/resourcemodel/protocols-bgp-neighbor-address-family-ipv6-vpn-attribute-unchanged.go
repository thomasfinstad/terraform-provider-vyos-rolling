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

var _ helpers.VyosResourceDataModel = &ProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged{}

// ProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedAsPath  types.Bool `tfsdk:"as_path" vyos:"as-path,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedMed     types.Bool `tfsdk:"med" vyos:"med,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchangedNextHop types.Bool `tfsdk:"next_hop" vyos:"next-hop,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
