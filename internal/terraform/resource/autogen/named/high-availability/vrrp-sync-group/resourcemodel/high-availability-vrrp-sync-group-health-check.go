// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &HighAvailabilityVrrpSyncGroupHealthCheck{}

// HighAvailabilityVrrpSyncGroupHealthCheck describes the resource data model.
type HighAvailabilityVrrpSyncGroupHealthCheck struct {
	// LeafNodes
	LeafHighAvailabilityVrrpSyncGroupHealthCheckFailureCount types.String `tfsdk:"failure_count" vyos:"failure-count,omitempty"`
	LeafHighAvailabilityVrrpSyncGroupHealthCheckInterval     types.String `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafHighAvailabilityVrrpSyncGroupHealthCheckPing         types.String `tfsdk:"ping" vyos:"ping,omitempty"`
	LeafHighAvailabilityVrrpSyncGroupHealthCheckScrIPt       types.String `tfsdk:"script" vyos:"script,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpSyncGroupHealthCheck) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"failure_count": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check failure count required for transition to fault

`,
			Description: `Health check failure count required for transition to fault

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check execution interval in seconds

`,
			Description: `Health check execution interval in seconds

`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"ping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMP ping health check

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  IPv4 ping target address  |
    |  ipv6    |  IPv6 ping target address  |
`,
			Description: `ICMP ping health check

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  IPv4 ping target address  |
    |  ipv6    |  IPv6 ping target address  |
`,
		},

		"script": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check script file

`,
			Description: `Health check script file

`,
		},

		// Nodes

	}
}
