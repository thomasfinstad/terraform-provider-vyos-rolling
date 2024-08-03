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
var _ helpers.VyosTopResourceDataModel = &VpnPptpRemoteAccessLog{}

// VpnPptpRemoteAccessLog describes the resource data model.
type VpnPptpRemoteAccessLog struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnPptpRemoteAccessLogLevel types.String `tfsdk:"level" vyos:"level,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *VpnPptpRemoteAccessLog) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnPptpRemoteAccessLog) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnPptpRemoteAccessLog) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnPptpRemoteAccessLog) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"log",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnPptpRemoteAccessLog) GetVyosParentPath() []string {
	return []string{
		"vpn",

		"pptp",

		"remote-access",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *VpnPptpRemoteAccessLog) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnPptpRemoteAccessLog) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies log level

    |  Format  |  Description                                          |
    |----------|-------------------------------------------------------|
    |  0       |  Turn off logging                                     |
    |  1       |  Log only error messages                              |
    |  2       |  Log error and warning messages                       |
    |  3       |  Log error, warning and minimum information messages  |
    |  4       |  Log error, warning and full information messages     |
    |  5       |  Log all messages including debug messages            |
`,
			Description: `Specifies log level

    |  Format  |  Description                                          |
    |----------|-------------------------------------------------------|
    |  0       |  Turn off logging                                     |
    |  1       |  Log only error messages                              |
    |  2       |  Log error and warning messages                       |
    |  3       |  Log error, warning and minimum information messages  |
    |  4       |  Log error, warning and full information messages     |
    |  5       |  Log all messages including debug messages            |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},
	}
}
