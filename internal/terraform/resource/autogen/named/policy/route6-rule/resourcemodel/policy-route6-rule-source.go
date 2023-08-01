// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRoutesixRuleSource describes the resource data model.
type PolicyRoutesixRuleSource struct {
	// LeafNodes
	LeafPolicyRoutesixRuleSourceAddress    types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafPolicyRoutesixRuleSourcePort       types.String `tfsdk:"port" vyos:"port,omitempty"`
	LeafPolicyRoutesixRuleSourceMacAddress types.String `tfsdk:"mac_address" vyos:"mac-address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyRoutesixRuleSourceGroup *PolicyRoutesixRuleSourceGroup `tfsdk:"group" vyos:"group,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6  &emsp; |  IP address to match  |
    |  ipv6net  &emsp; |  Subnet to match  |
    |  ipv6range  &emsp; |  IP range to match  |
    |  !ipv6  &emsp; |  Match everything except the specified address  |
    |  !ipv6net  &emsp; |  Match everything except the specified prefix  |
    |  !ipv6range  &emsp; |  Match everything except the specified range  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Named port (any name in /etc/services, e.g., http)  |
    |  number: 1-65535  &emsp; |  Numbered port  |
    |  <start-end>  &emsp; |  Numbered port range (e.g. 1001-1005)  |
    |     &emsp; |  \n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: 'telnet,http,123,1001-1005'  |

`,
		},

		"mac_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MAC address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macaddr  &emsp; |  MAC address to match  |
    |  !macaddr  &emsp; |  Match everything except the specified MAC address  |

`,
		},

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleSourceGroup{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Group

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRoutesixRuleSource) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRoutesixRuleSource) UnmarshalJSON(_ []byte) error {
	return nil
}
