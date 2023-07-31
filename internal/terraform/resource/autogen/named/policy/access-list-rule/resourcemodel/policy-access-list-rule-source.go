// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyAccessListRuleSource describes the resource data model.
type PolicyAccessListRuleSource struct {
	// LeafNodes
	LeafPolicyAccessListRuleSourceAny         types.Bool   `tfsdk:"any" vyos:"any,omitempty"`
	LeafPolicyAccessListRuleSourceHost        types.String `tfsdk:"host" vyos:"host,omitempty"`
	LeafPolicyAccessListRuleSourceInverseMask types.String `tfsdk:"inverse_mask" vyos:"inverse-mask,omitempty"`
	LeafPolicyAccessListRuleSourceNetwork     types.String `tfsdk:"network" vyos:"network,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessListRuleSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"any": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Any IP address to match

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"host": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Single host IP address to match

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  Host address to match  |

`,
		},

		"inverse_mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network/netmask to match (requires network be defined)

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  Inverse-mask to match  |

`,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Network/netmask to match (requires inverse-mask be defined)

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  Inverse-mask to match  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyAccessListRuleSource) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyAccessListRuleSource) UnmarshalJSON(_ []byte) error {
	return nil
}
