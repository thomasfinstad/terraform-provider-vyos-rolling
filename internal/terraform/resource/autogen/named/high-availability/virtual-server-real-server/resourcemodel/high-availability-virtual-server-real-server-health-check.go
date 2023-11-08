// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HighAvailabilityVirtualServerRealServerHealthCheck describes the resource data model.
type HighAvailabilityVirtualServerRealServerHealthCheck struct {
	// LeafNodes
	LeafHighAvailabilityVirtualServerRealServerHealthCheckScrIPt types.String `tfsdk:"script" vyos:"script,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVirtualServerRealServerHealthCheck) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"script": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health check script file

`,
		},

		// Nodes

	}
}
