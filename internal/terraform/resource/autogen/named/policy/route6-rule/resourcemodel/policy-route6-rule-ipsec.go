// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRoutesixRuleIPsec describes the resource data model.
type PolicyRoutesixRuleIPsec struct {
	// LeafNodes
	LeafPolicyRoutesixRuleIPsecMatchIPsec types.Bool `tfsdk:"match_ipsec" vyos:"match-ipsec,omitempty"`
	LeafPolicyRoutesixRuleIPsecMatchNone  types.Bool `tfsdk:"match_none" vyos:"match-none,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleIPsec) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_ipsec": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound IPsec packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"match_none": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inbound non-IPsec packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRoutesixRuleIPsec) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRoutesixRuleIPsec) UnmarshalJSON(_ []byte) error {
	return nil
}
