// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpPeerGroupBfd describes the resource data model.
type ProtocolsBgpPeerGroupBfd struct {
	// LeafNodes
	ProtocolsBgpPeerGroupBfdProfile                  customtypes.CustomStringValue `tfsdk:"profile" json:"profile,omitempty"`
	ProtocolsBgpPeerGroupBfdCheckControlPlaneFailure customtypes.CustomStringValue `tfsdk:"check_control_plane_failure" json:"check-control-plane-failure,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpPeerGroupBfd) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Use settings from BFD profile

|  Format  |  Description  |
|----------|---------------|
|  txt  |  BFD profile name  |

`,
		},

		"check_control_plane_failure": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Allow to write CBIT independence in BFD outgoing packets and read both C-BIT value of BFD and lookup BGP peer status

`,
		},

		// TagNodes

		// Nodes

	}
}
