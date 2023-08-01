// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRoutesixRuleSet describes the resource data model.
type PolicyRoutesixRuleSet struct {
	// LeafNodes
	LeafPolicyRoutesixRuleSetConnectionMark types.Number `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafPolicyRoutesixRuleSetDscp           types.Number `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafPolicyRoutesixRuleSetMark           types.Number `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafPolicyRoutesixRuleSetTable          types.String `tfsdk:"table" vyos:"table,omitempty"`
	LeafPolicyRoutesixRuleSetTCPMss         types.Number `tfsdk:"tcp_mss" vyos:"tcp-mss,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleSet) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"connection_mark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Connection marking

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-2147483647  &emsp; |  Connection marking  |

`,
		},

		"dscp": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Packet Differentiated Services Codepoint (DSCP)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-63  &emsp; |  DSCP number  |

`,
		},

		"mark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Packet marking

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2147483647  &emsp; |  Packet marking  |

`,
		},

		"table": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Routing table to forward packet with

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-200  &emsp; |  Table number  |
    |  main  &emsp; |  Main table  |

`,
		},

		"tcp_mss": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP Maximum Segment Size

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 500-1460  &emsp; |  Explicitly set TCP MSS value  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRoutesixRuleSet) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRoutesixRuleSet) UnmarshalJSON(_ []byte) error {
	return nil
}
