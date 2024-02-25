// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvfourInputFilterRuleIcmp describes the resource data model.
type FirewallIPvfourInputFilterRuleIcmp struct {
	// LeafNodes
	LeafFirewallIPvfourInputFilterRuleIcmpCode     types.Number `tfsdk:"code" vyos:"code,omitempty"`
	LeafFirewallIPvfourInputFilterRuleIcmpType     types.Number `tfsdk:"type" vyos:"type,omitempty"`
	LeafFirewallIPvfourInputFilterRuleIcmpTypeName types.String `tfsdk:"type_name" vyos:"type-name,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourInputFilterRuleIcmp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"code": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ICMP code

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-255  &emsp; |  ICMP code (0-255)  |

`,
		},

		"type": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ICMP type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-255  &emsp; |  ICMP type (0-255)  |

`,
		},

		"type_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMP type-name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  echo-reply  &emsp; |  ICMP type 0: echo-reply  |
    |  destination-unreachable  &emsp; |  ICMP type 3: destination-unreachable  |
    |  source-quench  &emsp; |  ICMP type 4: source-quench  |
    |  redirect  &emsp; |  ICMP type 5: redirect  |
    |  echo-request  &emsp; |  ICMP type 8: echo-request  |
    |  router-advertisement  &emsp; |  ICMP type 9: router-advertisement  |
    |  router-solicitation  &emsp; |  ICMP type 10: router-solicitation  |
    |  time-exceeded  &emsp; |  ICMP type 11: time-exceeded  |
    |  parameter-problem  &emsp; |  ICMP type 12: parameter-problem  |
    |  timestamp-request  &emsp; |  ICMP type 13: timestamp-request  |
    |  timestamp-reply  &emsp; |  ICMP type 14: timestamp-reply  |
    |  info-request  &emsp; |  ICMP type 15: info-request  |
    |  info-reply  &emsp; |  ICMP type 16: info-reply  |
    |  address-mask-request  &emsp; |  ICMP type 17: address-mask-request  |
    |  address-mask-reply  &emsp; |  ICMP type 18: address-mask-reply  |

`,
		},

		// Nodes

	}
}
