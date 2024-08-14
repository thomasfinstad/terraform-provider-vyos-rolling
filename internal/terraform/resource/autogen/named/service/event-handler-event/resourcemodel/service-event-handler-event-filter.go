// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceEventHandlerEventFilter{}

// ServiceEventHandlerEventFilter describes the resource data model.
type ServiceEventHandlerEventFilter struct {
	// LeafNodes
	LeafServiceEventHandlerEventFilterPattern          types.String `tfsdk:"pattern" vyos:"pattern,omitempty"`
	LeafServiceEventHandlerEventFilterSyslogIDentifier types.String `tfsdk:"syslog_identifier" vyos:"syslog-identifier,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceEventHandlerEventFilter) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"pattern": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match pattern (regex)

`,
			Description: `Match pattern (regex)

`,
		},

		"syslog_identifier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identifier of a process in syslog (string)

`,
			Description: `Identifier of a process in syslog (string)

`,
		},

		// Nodes

	}
}