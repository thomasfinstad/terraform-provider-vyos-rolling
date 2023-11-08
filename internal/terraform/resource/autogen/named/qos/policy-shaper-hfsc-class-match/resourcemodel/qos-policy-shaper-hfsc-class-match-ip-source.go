// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyShaperHfscClassMatchIPSource describes the resource data model.
type QosPolicyShaperHfscClassMatchIPSource struct {
	// LeafNodes
	LeafQosPolicyShaperHfscClassMatchIPSourceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafQosPolicyShaperHfscClassMatchIPSourcePort    types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatchIPSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 destination address for this match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address  |
    |  ipv4net  &emsp; |  IPv4 prefix  |

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,
		},

		// Nodes

	}
}
