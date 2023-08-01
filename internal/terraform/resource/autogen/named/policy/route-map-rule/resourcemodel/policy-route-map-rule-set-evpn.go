// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// PolicyRouteMapRuleSetEvpn describes the resource data model.
type PolicyRouteMapRuleSetEvpn struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyRouteMapRuleSetEvpnGateway *PolicyRouteMapRuleSetEvpnGateway `tfsdk:"gateway" vyos:"gateway,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetEvpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"gateway": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetEvpnGateway{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Set gateway IP for prefix advertisement route

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRuleSetEvpn) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRouteMapRuleSetEvpn) UnmarshalJSON(_ []byte) error {
	return nil
}