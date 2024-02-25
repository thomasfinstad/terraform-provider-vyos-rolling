// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvsixOutputFilterRuleLogOptions describes the resource data model.
type FirewallIPvsixOutputFilterRuleLogOptions struct {
	// LeafNodes
	LeafFirewallIPvsixOutputFilterRuleLogOptionsGroup          types.Number `tfsdk:"group" vyos:"group,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleLogOptionsSnapshotLength types.Number `tfsdk:"snapshot_length" vyos:"snapshot-length,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleLogOptionsQueueThreshold types.Number `tfsdk:"queue_threshold" vyos:"queue-threshold,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleLogOptionsLevel          types.String `tfsdk:"level" vyos:"level,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputFilterRuleLogOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"group": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set log group

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Log group to send messages to  |

`,
		},

		"snapshot_length": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Length of packet payload to include in netlink message

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-9000  &emsp; |  Length of packet payload to include in netlink message  |

`,
		},

		"queue_threshold": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of packets to queue inside the kernel before sending them to userspace

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Number of packets to queue inside the kernel before sending them to userspace  |

`,
		},

		"level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set log-level

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  emerg  &emsp; |  Emerg log level  |
    |  alert  &emsp; |  Alert log level  |
    |  crit  &emsp; |  Critical log level  |
    |  err  &emsp; |  Error log level  |
    |  warn  &emsp; |  Warning log level  |
    |  notice  &emsp; |  Notice log level  |
    |  info  &emsp; |  Info log level  |
    |  debug  &emsp; |  Debug log level  |

`,
		},

		// Nodes

	}
}
