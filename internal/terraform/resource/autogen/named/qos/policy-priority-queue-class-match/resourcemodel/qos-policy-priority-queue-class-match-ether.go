// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyPriorityQueueClassMatchEther describes the resource data model.
type QosPolicyPriorityQueueClassMatchEther struct {
	// LeafNodes
	LeafQosPolicyPriorityQueueClassMatchEtherDestination types.String `tfsdk:"destination" vyos:"destination,omitempty"`
	LeafQosPolicyPriorityQueueClassMatchEtherProtocol    types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafQosPolicyPriorityQueueClassMatchEtherSource      types.String `tfsdk:"source" vyos:"source,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyPriorityQueueClassMatchEther) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet destination address for this match

    |  Format  |  Description  |
    |----------|---------------|
    |  macaddr  |  MAC address to match  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet protocol for this match

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-65535  |  Ethernet protocol number  |
    |  txt  |  Ethernet protocol name  |
    |  all  |  Any protocol  |
    |  ip  |  Internet IP (IPv4)  |
    |  ipv6  |  Internet IP (IPv6)  |
    |  arp  |  Address Resolution Protocol  |
    |  atalk  |  Appletalk  |
    |  ipx  |  Novell Internet Packet Exchange  |
    |  802.1Q  |  802.1Q VLAN tag  |

`,
		},

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet source address for this match

    |  Format  |  Description  |
    |----------|---------------|
    |  macaddr  |  MAC address to match  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyPriorityQueueClassMatchEther) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyPriorityQueueClassMatchEther) UnmarshalJSON(_ []byte) error {
	return nil
}
