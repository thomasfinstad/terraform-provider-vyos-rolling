// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRouteRuleTCP describes the resource data model.
type PolicyRouteRuleTCP struct {
	// LeafNodes
	LeafPolicyRouteRuleTCPMss types.String `tfsdk:"mss" vyos:"mss,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyRouteRuleTCPFlags *PolicyRouteRuleTCPFlags `tfsdk:"flags" vyos:"flags,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum segment size (MSS)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-16384  |  Maximum segment size  |
    |  <min>-<max>  |  TCP MSS range (use '-' as delimiter)  |

`,
		},

		// Nodes

		"flags": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleTCPFlags{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteRuleTCP) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRouteRuleTCP) UnmarshalJSON(_ []byte) error {
	return nil
}
