// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyLimiterClassMatchIPvsixDestination describes the resource data model.
type QosPolicyLimiterClassMatchIPvsixDestination struct {
	// LeafNodes
	LeafQosPolicyLimiterClassMatchIPvsixDestinationAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafQosPolicyLimiterClassMatchIPvsixDestinationPort    types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyLimiterClassMatchIPvsixDestination) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 destination address for this match

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv6net  |  IPv6 address and prefix length  |

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Numeric IP port  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyLimiterClassMatchIPvsixDestination) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyLimiterClassMatchIPvsixDestination) UnmarshalJSON(_ []byte) error {
	return nil
}
