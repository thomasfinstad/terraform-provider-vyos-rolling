// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRouteMapRuleMatchIPvsixAddress describes the resource data model.
type PolicyRouteMapRuleMatchIPvsixAddress struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPvsixAddressAccessList types.String `tfsdk:"access_list" vyos:"access-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixAddressPrefixList types.String `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixAddressPrefixLen  types.Number `tfsdk:"prefix_len" vyos:"prefix-len,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPvsixAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"access_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 access-list to match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  IPV6 access list name  |

`,
		},

		"prefix_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 prefix-list to match

`,
		},

		"prefix_len": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 prefix-length to match (can be used for kernel routes only)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-128  &emsp; |  Prefix length  |

`,
		},

		// Nodes

	}
}
