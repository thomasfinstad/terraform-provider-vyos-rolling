// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ServiceIPoeServerExtendedScrIPts{}

// ServiceIPoeServerExtendedScrIPts describes the resource data model.
type ServiceIPoeServerExtendedScrIPts struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceIPoeServerExtendedScrIPtsOnPreUp  types.String `tfsdk:"on_pre_up" vyos:"on-pre-up,omitempty"`
	LeafServiceIPoeServerExtendedScrIPtsOnUp     types.String `tfsdk:"on_up" vyos:"on-up,omitempty"`
	LeafServiceIPoeServerExtendedScrIPtsOnDown   types.String `tfsdk:"on_down" vyos:"on-down,omitempty"`
	LeafServiceIPoeServerExtendedScrIPtsOnChange types.String `tfsdk:"on_change" vyos:"on-change,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceIPoeServerExtendedScrIPts) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceIPoeServerExtendedScrIPts) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceIPoeServerExtendedScrIPts) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerExtendedScrIPts) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"extended-scripts",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceIPoeServerExtendedScrIPts) GetVyosParentPath() []string {
	return []string{
		"service",

		"ipoe-server",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ServiceIPoeServerExtendedScrIPts) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerExtendedScrIPts) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"on_pre_up": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run before session interface comes up

`,
			Description: `Script to run before session interface comes up

`,
		},

		"on_up": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run when session interface is completely configured and started

`,
			Description: `Script to run when session interface is completely configured and started

`,
		},

		"on_down": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run when session interface going to terminate

`,
			Description: `Script to run when session interface going to terminate

`,
		},

		"on_change": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run when session interface changed by RADIUS CoA handling

`,
			Description: `Script to run when session interface changed by RADIUS CoA handling

`,
		},
	}
}
