// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyLocalRoutesixRule describes the resource data model.
type PolicyLocalRoutesixRule struct {
	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:",self-id"`

	// LeafNodes
	LeafPolicyLocalRoutesixRuleFwmark           types.Number `tfsdk:"fwmark" vyos:"fwmark,omitempty"`
	LeafPolicyLocalRoutesixRuleSource           types.List   `tfsdk:"source" vyos:"source,omitempty"`
	LeafPolicyLocalRoutesixRuleDestination      types.List   `tfsdk:"destination" vyos:"destination,omitempty"`
	LeafPolicyLocalRoutesixRuleInboundInterface types.String `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyLocalRoutesixRuleSet *PolicyLocalRoutesixRuleSet `tfsdk:"set" vyos:"set,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyLocalRoutesixRule) GetVyosPath() []string {
	return []string{
		"policy",

		"local-route6",

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRoutesixRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 policy local-route rule set number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-32765  &emsp; |  Local-route rule number (1-32765)  |

`,
		},

		// LeafNodes

		"fwmark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match fwmark value

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2147483647  &emsp; |  Address to match against  |

`,
		},

		"source": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Source address or prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6  &emsp; |  Address to match against  |
    |  ipv6net  &emsp; |  Prefix to match against  |

`,
		},

		"destination": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Destination address or prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6  &emsp; |  Address to match against  |
    |  ipv6net  &emsp; |  Prefix to match against  |

`,
		},

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound Interface

`,
		},

		// Nodes

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyLocalRoutesixRuleSet{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyLocalRoutesixRule) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyLocalRoutesixRule) UnmarshalJSON(_ []byte) error {
	return nil
}
