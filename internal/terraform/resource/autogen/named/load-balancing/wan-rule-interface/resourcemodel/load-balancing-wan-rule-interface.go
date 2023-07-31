// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// LoadBalancingWanRuleInterface describes the resource data model.
type LoadBalancingWanRuleInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDLoadBalancingWanRule types.String `tfsdk:"rule" vyos:"rule_identifier,parent-id"`

	// LeafNodes
	LeafLoadBalancingWanRuleInterfaceWeight types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingWanRuleInterface) GetVyosPath() []string {
	return []string{
		"load-balancing",

		"wan",

		"rule",
		o.ParentIDLoadBalancingWanRule.ValueString(),

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingWanRuleInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface name [REQUIRED]

`,
		},

		"rule_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule number (1-9999)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-9999  |  Rule number  |

`,
		},

		// LeafNodes

		"weight": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Load-balance weight

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Interface weight  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *LoadBalancingWanRuleInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *LoadBalancingWanRuleInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
