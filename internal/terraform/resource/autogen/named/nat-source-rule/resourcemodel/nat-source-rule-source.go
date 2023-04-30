// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// NatSourceRuleSource describes the resource data model.
type NatSourceRuleSource struct {
	// LeafNodes
	NatSourceRuleSourceAddress customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	NatSourceRuleSourcePort    customtypes.CustomStringValue `tfsdk:"port" json:"port,omitempty"`

	// TagNodes

	// Nodes
	NatSourceRuleSourceGroup types.Object `tfsdk:"group" json:"group,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o NatSourceRuleSource) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |
|  ipv4range  |  IPv4 address range to match  |
|  !ipv4  |  Match everything except the specified address  |
|  !ipv4net  |  Match everything except the specified prefix  |
|  !ipv4range  |  Match everything except the specified range  |

`,
		},

		"port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Port number

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numeric IP port  |
|  start-end  |  Numbered port range (e.g. 1001-1005)  |
|   |   |

`,
		},

		// TagNodes

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: NatSourceRuleSourceGroup{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Group

`,
		},
	}
}
