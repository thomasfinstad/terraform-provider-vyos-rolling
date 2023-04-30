// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRouteMapRuleSetEvpnGateway describes the resource data model.
type PolicyRouteMapRuleSetEvpnGateway struct {
	// LeafNodes
	PolicyRouteMapRuleSetEvpnGatewayIPvfour customtypes.CustomStringValue `tfsdk:"ipv4" json:"ipv4,omitempty"`
	PolicyRouteMapRuleSetEvpnGatewayIPvsix  customtypes.CustomStringValue `tfsdk:"ipv6" json:"ipv6,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRouteMapRuleSetEvpnGateway) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ipv4": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Set gateway IPv4 address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Gateway IPv4 address  |

`,
		},

		"ipv6": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Set gateway IPv6 address

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Gateway IPv6 address  |

`,
		},

		// TagNodes

		// Nodes

	}
}
