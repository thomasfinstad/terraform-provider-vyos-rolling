/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (log-options) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgePreroutingFilterRuleLogOptions{}

// FirewallBrIDgePreroutingFilterRuleLogOptions describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallBrIDgePreroutingFilterRuleLogOptions struct {
	// LeafNodes
	LeafFirewallBrIDgePreroutingFilterRuleLogOptionsGroup          types.Number `tfsdk:"group" vyos:"group,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleLogOptionsSnapshotLength types.Number `tfsdk:"snapshot_length" vyos:"snapshot-length,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleLogOptionsQueueThreshold types.Number `tfsdk:"queue_threshold" vyos:"queue-threshold,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleLogOptionsLevel          types.String `tfsdk:"level" vyos:"level,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgePreroutingFilterRuleLogOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (group) */
		schema.NumberAttribute{
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

		"snapshot_length":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (snapshot-length) */
		schema.NumberAttribute{
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

		"queue_threshold":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (queue-threshold) */
		schema.NumberAttribute{
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

		"level":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (level) */
		schema.StringAttribute{
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

		// TagNodes

		// Nodes

	}
}
