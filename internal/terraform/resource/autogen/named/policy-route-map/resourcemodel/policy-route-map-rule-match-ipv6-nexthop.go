// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRouteMapRuleMatchIPvsixNexthop describes the resource data model.
type PolicyRouteMapRuleMatchIPvsixNexthop struct {
	// LeafNodes
	PolicyRouteMapRuleMatchIPvsixNexthopAddress    customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	PolicyRouteMapRuleMatchIPvsixNexthopAccessList customtypes.CustomStringValue `tfsdk:"access_list" json:"access-list,omitempty"`
	PolicyRouteMapRuleMatchIPvsixNexthopPrefixList customtypes.CustomStringValue `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
	PolicyRouteMapRuleMatchIPvsixNexthopType       customtypes.CustomStringValue `tfsdk:"type" json:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPvsixNexthop) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv6 address of next-hop

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Nexthop IPv6 address  |

`,
		},

		"access_list": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv6 access-list to match

|  Format  |  Description  |
|----------|---------------|
|  txt  |  IPV6 access list name  |

`,
		},

		"prefix_list": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv6 prefix-list to match

`,
		},

		"type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Match type

|  Format  |  Description  |
|----------|---------------|
|  blackhole  |  Blackhole  |

`,
		},

		// TagNodes

		// Nodes

	}
}
