// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpPeerGroupBfd describes the resource data model.
type VrfNameProtocolsBgpPeerGroupBfd struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupBfdProfile                  types.String `tfsdk:"profile" vyos:"profile,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupBfdCheckControlPlaneFailure types.Bool   `tfsdk:"check_control_plane_failure" vyos:"check-control-plane-failure,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupBfd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  BFD profile name  |

`,
		},

		"check_control_plane_failure": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allow to write CBIT independence in BFD outgoing packets and read both C-BIT value of BFD and lookup BGP peer status

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpPeerGroupBfd) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpPeerGroupBfd) UnmarshalJSON(_ []byte) error {
	return nil
}
