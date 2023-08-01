// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallZoneFrom describes the resource data model.
type FirewallZoneFrom struct {
	SelfIdentifier types.String `tfsdk:"from_id" vyos:",self-id"`

	ParentIDFirewallZone types.String `tfsdk:"zone" vyos:"zone,parent-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallZoneFromFirewall *FirewallZoneFromFirewall `tfsdk:"firewall" vyos:"firewall,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallZoneFrom) GetVyosPath() []string {
	return []string{
		"firewall",

		"zone",
		o.ParentIDFirewallZone.ValueString(),

		"from",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZoneFrom) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"from_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Zone from which to filter traffic

`,
		},

		"zone_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Zone-policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Zone name  |

`,
		},

		// LeafNodes

		// Nodes

		"firewall": schema.SingleNestedAttribute{
			Attributes: FirewallZoneFromFirewall{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Firewall options

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallZoneFrom) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallZoneFrom) UnmarshalJSON(_ []byte) error {
	return nil
}