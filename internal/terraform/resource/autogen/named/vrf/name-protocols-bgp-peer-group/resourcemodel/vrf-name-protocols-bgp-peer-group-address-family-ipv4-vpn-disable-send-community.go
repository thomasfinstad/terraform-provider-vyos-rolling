// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnDisableSendCommunity describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnDisableSendCommunity struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnDisableSendCommunityExtended types.Bool `tfsdk:"extended" vyos:"extended,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnDisableSendCommunityStandard types.Bool `tfsdk:"standard" vyos:"standard,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnDisableSendCommunity) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"extended": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable sending extended community attributes to this peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"standard": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable sending standard community attributes to this peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
