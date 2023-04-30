// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsStaticRouteInterface describes the resource data model.
type ProtocolsStaticRouteInterface struct {
	// LeafNodes
	ProtocolsStaticRouteInterfaceDisable  customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	ProtocolsStaticRouteInterfaceDistance customtypes.CustomStringValue `tfsdk:"distance" json:"distance,omitempty"`
	ProtocolsStaticRouteInterfaceVrf      customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsStaticRouteInterface) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable instance

`,
		},

		"distance": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Distance for this route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for this route  |

`,
		},

		"vrf": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VRF to leak route

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of VRF to leak to  |

`,
		},

		// TagNodes

		// Nodes

	}
}
