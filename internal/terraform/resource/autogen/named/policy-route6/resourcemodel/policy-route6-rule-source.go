// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRoutesixRuleSource describes the resource data model.
type PolicyRoutesixRuleSource struct {
	// LeafNodes
	PolicyRoutesixRuleSourceAddress    customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	PolicyRoutesixRuleSourcePort       customtypes.CustomStringValue `tfsdk:"port" json:"port,omitempty"`
	PolicyRoutesixRuleSourceMacAddress customtypes.CustomStringValue `tfsdk:"mac_address" json:"mac-address,omitempty"`

	// TagNodes

	// Nodes
	PolicyRoutesixRuleSourceGroup types.Object `tfsdk:"group" json:"group,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRoutesixRuleSource) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IP address to match  |
|  ipv6net  |  Subnet to match  |
|  ipv6range  |  IP range to match  |
|  !ipv6  |  Match everything except the specified address  |
|  !ipv6net  |  Match everything except the specified prefix  |
|  !ipv6range  |  Match everything except the specified range  |

`,
		},

		"port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Port

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numbered port  |
|  <start-end>  |  Numbered port range (e.g. 1001-1005)  |
|     |  \n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: 'telnet,http,123,1001-1005'  |

`,
		},

		"mac_address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `MAC address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  MAC address to match  |
|  !macaddr  |  Match everything except the specified MAC address  |

`,
		},

		// TagNodes

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleSourceGroup{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Group

`,
		},
	}
}
