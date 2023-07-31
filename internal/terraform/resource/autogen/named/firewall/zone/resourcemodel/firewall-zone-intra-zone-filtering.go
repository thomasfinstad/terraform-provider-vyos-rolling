// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallZoneIntraZoneFiltering describes the resource data model.
type FirewallZoneIntraZoneFiltering struct {
	// LeafNodes
	LeafFirewallZoneIntraZoneFilteringAction types.String `tfsdk:"action" vyos:"action,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallZoneIntraZoneFilteringFirewall *FirewallZoneIntraZoneFilteringFirewall `tfsdk:"firewall" vyos:"firewall,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZoneIntraZoneFiltering) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action for intra-zone traffic

    |  Format  |  Description  |
    |----------|---------------|
    |  accept  |  Accept traffic  |
    |  drop  |  Drop silently  |

`,
		},

		// Nodes

		"firewall": schema.SingleNestedAttribute{
			Attributes: FirewallZoneIntraZoneFilteringFirewall{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use the specified firewall chain

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallZoneIntraZoneFiltering) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallZoneIntraZoneFiltering) UnmarshalJSON(_ []byte) error {
	return nil
}
