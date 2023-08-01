// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NatStaticRuleTranSLAtion describes the resource data model.
type NatStaticRuleTranSLAtion struct {
	// LeafNodes
	LeafNatStaticRuleTranSLAtionAddress types.String `tfsdk:"address" vyos:"address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatStaticRuleTranSLAtion) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to match  |
    |  ipv4net  &emsp; |  IPv4 prefix to match  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *NatStaticRuleTranSLAtion) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *NatStaticRuleTranSLAtion) UnmarshalJSON(_ []byte) error {
	return nil
}
