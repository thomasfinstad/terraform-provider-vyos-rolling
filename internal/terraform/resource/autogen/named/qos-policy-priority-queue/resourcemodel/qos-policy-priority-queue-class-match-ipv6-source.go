// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyPriorityQueueClassMatchIPvsixSource describes the resource data model.
type QosPolicyPriorityQueueClassMatchIPvsixSource struct {
	// LeafNodes
	QosPolicyPriorityQueueClassMatchIPvsixSourceAddress customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	QosPolicyPriorityQueueClassMatchIPvsixSourcePort    customtypes.CustomStringValue `tfsdk:"port" json:"port,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyPriorityQueueClassMatchIPvsixSource) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv6 destination address for this match

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |

`,
		},

		"port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,
		},

		// TagNodes

		// Nodes

	}
}
