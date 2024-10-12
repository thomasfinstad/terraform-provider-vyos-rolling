// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputRawRuleLogOptions{}

// FirewallIPvfourOutputRawRuleLogOptions describes the resource data model.
type FirewallIPvfourOutputRawRuleLogOptions struct {
	// LeafNodes
	LeafFirewallIPvfourOutputRawRuleLogOptionsGroup          types.Number `tfsdk:"group" vyos:"group,omitempty"`
	LeafFirewallIPvfourOutputRawRuleLogOptionsSnapshotLength types.Number `tfsdk:"snapshot_length" vyos:"snapshot-length,omitempty"`
	LeafFirewallIPvfourOutputRawRuleLogOptionsQueueThreshold types.Number `tfsdk:"queue_threshold" vyos:"queue-threshold,omitempty"`
	LeafFirewallIPvfourOutputRawRuleLogOptionsLevel          types.String `tfsdk:"level" vyos:"level,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputRawRuleLogOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"group": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set log group

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  0-65535  |  Log group to send messages to  |
`,
			Description: `Set log group

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  0-65535  |  Log group to send messages to  |
`,
		},

		"snapshot_length": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Length of packet payload to include in netlink message

    |  Format  |  Description                                             |
    |----------|----------------------------------------------------------|
    |  0-9000  |  Length of packet payload to include in netlink message  |
`,
			Description: `Length of packet payload to include in netlink message

    |  Format  |  Description                                             |
    |----------|----------------------------------------------------------|
    |  0-9000  |  Length of packet payload to include in netlink message  |
`,
		},

		"queue_threshold": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of packets to queue inside the kernel before sending them to userspace

    |  Format   |  Description                                                                    |
    |-----------|---------------------------------------------------------------------------------|
    |  0-65535  |  Number of packets to queue inside the kernel before sending them to userspace  |
`,
			Description: `Number of packets to queue inside the kernel before sending them to userspace

    |  Format   |  Description                                                                    |
    |-----------|---------------------------------------------------------------------------------|
    |  0-65535  |  Number of packets to queue inside the kernel before sending them to userspace  |
`,
		},

		"level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set log-level

    |  Format  |  Description         |
    |----------|----------------------|
    |  emerg   |  Emerg log level     |
    |  alert   |  Alert log level     |
    |  crit    |  Critical log level  |
    |  err     |  Error log level     |
    |  warn    |  Warning log level   |
    |  notice  |  Notice log level    |
    |  info    |  Info log level      |
    |  debug   |  Debug log level     |
`,
			Description: `Set log-level

    |  Format  |  Description         |
    |----------|----------------------|
    |  emerg   |  Emerg log level     |
    |  alert   |  Alert log level     |
    |  crit    |  Critical log level  |
    |  err     |  Error log level     |
    |  warn    |  Warning log level   |
    |  notice  |  Notice log level    |
    |  info    |  Info log level      |
    |  debug   |  Debug log level     |
`,
		},

		// Nodes

	}
}
