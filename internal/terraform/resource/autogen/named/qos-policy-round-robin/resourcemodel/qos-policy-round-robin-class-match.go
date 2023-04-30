// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyRoundRobinClassMatch describes the resource data model.
type QosPolicyRoundRobinClassMatch struct {
	// LeafNodes
	QosPolicyRoundRobinClassMatchDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	QosPolicyRoundRobinClassMatchInterface   customtypes.CustomStringValue `tfsdk:"interface" json:"interface,omitempty"`
	QosPolicyRoundRobinClassMatchMark        customtypes.CustomStringValue `tfsdk:"mark" json:"mark,omitempty"`
	QosPolicyRoundRobinClassMatchVif         customtypes.CustomStringValue `tfsdk:"vif" json:"vif,omitempty"`

	// TagNodes

	// Nodes
	QosPolicyRoundRobinClassMatchEther  types.Object `tfsdk:"ether" json:"ether,omitempty"`
	QosPolicyRoundRobinClassMatchIP     types.Object `tfsdk:"ip" json:"ip,omitempty"`
	QosPolicyRoundRobinClassMatchIPvsix types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyRoundRobinClassMatch) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"interface": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interface to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		},

		"mark": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Match on mark applied by firewall

|  Format  |  Description  |
|----------|---------------|
|  u32  |  FW mark to match  |

`,
		},

		"vif": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID for this match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4095  |  Virtual Local Area Network (VLAN) tag   |

`,
		},

		// TagNodes

		// Nodes

		"ether": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinClassMatchEther{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Ethernet header match

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinClassMatchIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IP protocol header

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinClassMatchIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IPv6 protocol header

`,
		},
	}
}
