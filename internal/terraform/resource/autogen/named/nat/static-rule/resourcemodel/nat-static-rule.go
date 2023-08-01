// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NatStaticRule describes the resource data model.
type NatStaticRule struct {
	SelfIdentifier types.String `tfsdk:"rule_id" vyos:",self-id"`

	// LeafNodes
	LeafNatStaticRuleDescrIPtion      types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafNatStaticRuleInboundInterface types.String `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeNatStaticRuleDestination *NatStaticRuleDestination `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeNatStaticRuleTranSLAtion *NatStaticRuleTranSLAtion `tfsdk:"translation" vyos:"translation,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *NatStaticRule) GetVyosPath() []string {
	return []string{
		"nat",

		"static",

		"rule",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatStaticRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule number for NAT

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound interface of NAT traffic

`,
		},

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: NatStaticRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `NAT destination parameters

`,
		},

		"translation": schema.SingleNestedAttribute{
			Attributes: NatStaticRuleTranSLAtion{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Translation address or prefix

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *NatStaticRule) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *NatStaticRule) UnmarshalJSON(_ []byte) error {
	return nil
}
