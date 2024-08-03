// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
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
var _ helpers.VyosTopResourceDataModel = &SystemLogsLogrotateAtop{}

// SystemLogsLogrotateAtop describes the resource data model.
type SystemLogsLogrotateAtop struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemLogsLogrotateAtopMaxSize types.Number `tfsdk:"max_size" vyos:"max-size,omitempty"`
	LeafSystemLogsLogrotateAtopRotate  types.Number `tfsdk:"rotate" vyos:"rotate,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemLogsLogrotateAtop) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemLogsLogrotateAtop) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemLogsLogrotateAtop) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemLogsLogrotateAtop) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"atop",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemLogsLogrotateAtop) GetVyosParentPath() []string {
	return []string{
		"system",

		"logs",

		"logrotate",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *SystemLogsLogrotateAtop) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLogsLogrotateAtop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"max_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Size of a single log file that triggers rotation

    |  Format  |  Description  |
    |----------|---------------|
    |  1-1024  |  Size in MB   |
`,
			Description: `Size of a single log file that triggers rotation

    |  Format  |  Description  |
    |----------|---------------|
    |  1-1024  |  Size in MB   |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"rotate": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Count of rotations before old logs will be deleted

    |  Format  |  Description  |
    |----------|---------------|
    |  1-100   |  Rotations    |
`,
			Description: `Count of rotations before old logs will be deleted

    |  Format  |  Description  |
    |----------|---------------|
    |  1-100   |  Rotations    |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},
	}
}
