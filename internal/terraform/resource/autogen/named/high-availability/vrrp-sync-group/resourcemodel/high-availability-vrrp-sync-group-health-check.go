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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (health-check) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &HighAvailabilityVrrpSyncGroupHealthCheck{}

// HighAvailabilityVrrpSyncGroupHealthCheck describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type HighAvailabilityVrrpSyncGroupHealthCheck struct {
	// LeafNodes
	LeafHighAvailabilityVrrpSyncGroupHealthCheckFailureCount types.String `tfsdk:"failure_count" vyos:"failure-count,omitempty"`
	LeafHighAvailabilityVrrpSyncGroupHealthCheckInterval     types.String `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafHighAvailabilityVrrpSyncGroupHealthCheckPing         types.String `tfsdk:"ping" vyos:"ping,omitempty"`
	LeafHighAvailabilityVrrpSyncGroupHealthCheckScrIPt       types.String `tfsdk:"script" vyos:"script,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpSyncGroupHealthCheck) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"failure_count":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (failure-count) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check failure count required for transition to fault

`,
			Description: `Health check failure count required for transition to fault

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (interval) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check execution interval in seconds

`,
			Description: `Health check execution interval in seconds

`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"ping":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ping) */
		schema.StringAttribute{
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

		"script":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (script) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check script file

`,
			Description: `Health check script file

`,
		},

		// TagNodes

		// Nodes

	}
}
