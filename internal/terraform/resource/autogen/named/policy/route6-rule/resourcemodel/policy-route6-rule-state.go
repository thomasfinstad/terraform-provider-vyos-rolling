// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRoutesixRuleState describes the resource data model.
type PolicyRoutesixRuleState struct {
	// LeafNodes
	LeafPolicyRoutesixRuleStateEstablished types.String `tfsdk:"established" vyos:"established,omitempty"`
	LeafPolicyRoutesixRuleStateInvalID     types.String `tfsdk:"invalid" vyos:"invalid,omitempty"`
	LeafPolicyRoutesixRuleStateNew         types.String `tfsdk:"new" vyos:"new,omitempty"`
	LeafPolicyRoutesixRuleStateRelated     types.String `tfsdk:"related" vyos:"related,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleState) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"established": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Established state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable  |
    |  disable  &emsp; |  Disable  |

`,
		},

		"invalid": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Invalid state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable  |
    |  disable  &emsp; |  Disable  |

`,
		},

		"new": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `New state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable  |
    |  disable  &emsp; |  Disable  |

`,
		},

		"related": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Related state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable  |
    |  disable  &emsp; |  Disable  |

`,
		},

		// Nodes

	}
}
