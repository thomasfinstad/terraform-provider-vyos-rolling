// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfigurationInbound types.Bool `tfsdk:"inbound" vyos:"inbound,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"inbound": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable inbound soft reconfiguration

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
