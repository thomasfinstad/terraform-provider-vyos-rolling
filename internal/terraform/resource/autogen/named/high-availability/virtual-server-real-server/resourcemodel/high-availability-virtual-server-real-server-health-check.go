// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &HighAvailabilityVirtualServerRealServerHealthCheck{}

// HighAvailabilityVirtualServerRealServerHealthCheck describes the resource data model.
type HighAvailabilityVirtualServerRealServerHealthCheck struct {
	// LeafNodes
	LeafHighAvailabilityVirtualServerRealServerHealthCheckScrIPt types.String `tfsdk:"script" vyos:"script,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVirtualServerRealServerHealthCheck) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

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
