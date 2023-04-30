// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// NatStaticRule describes the resource data model.
type NatStaticRule struct {
	// LeafNodes
	NatStaticRuleDescrIPtion      customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	NatStaticRuleInboundInterface customtypes.CustomStringValue `tfsdk:"inbound_interface" json:"inbound-interface,omitempty"`

	// TagNodes

	// Nodes
	NatStaticRuleDestination types.Object `tfsdk:"destination" json:"destination,omitempty"`
	NatStaticRuleTranSLAtion types.Object `tfsdk:"translation" json:"translation,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o NatStaticRule) ResourceAttributes() map[string]schema.Attribute {
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

		"inbound_interface": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Inbound interface of NAT traffic

`,
		},

		// TagNodes

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: NatStaticRuleDestination{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `NAT destination parameters

`,
		},

		"translation": schema.SingleNestedAttribute{
			Attributes: NatStaticRuleTranSLAtion{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Translation address or prefix

`,
		},
	}
}
