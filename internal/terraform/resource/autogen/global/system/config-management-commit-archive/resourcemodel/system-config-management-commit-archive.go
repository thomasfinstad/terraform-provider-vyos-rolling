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
var _ helpers.VyosTopResourceDataModel = &SystemConfigManagementCommitArchive{}

// SystemConfigManagementCommitArchive describes the resource data model.
type SystemConfigManagementCommitArchive struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemConfigManagementCommitArchiveLocation      types.List   `tfsdk:"location" vyos:"location,omitempty"`
	LeafSystemConfigManagementCommitArchiveSourceAddress types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemConfigManagementCommitArchive) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemConfigManagementCommitArchive) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemConfigManagementCommitArchive) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemConfigManagementCommitArchive) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"commit-archive",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemConfigManagementCommitArchive) GetVyosParentPath() []string {
	return []string{
		"system",

		"config-management",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *SystemConfigManagementCommitArchive) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConfigManagementCommitArchive) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"location": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Commit archive location

    |  Format                                     |  Description  |
    |---------------------------------------------|---------------|
    |  http://<user>:<passwd>@<host>/<path>       |  N/A          |
    |  https://<user>:<passwd>@<host>/<path>      |  N/A          |
    |  ftp://<user>:<passwd>@<host>/<path>        |  N/A          |
    |  sftp://<user>:<passwd>@<host>/<path>       |  N/A          |
    |  scp://<user>:<passwd>@<host>/<path>        |  N/A          |
    |  tftp://<host>/<path>                       |  N/A          |
    |  git+https://<user>:<passwd>@<host>/<path>  |  N/A          |
`,
			Description: `Commit archive location

    |  Format                                     |  Description  |
    |---------------------------------------------|---------------|
    |  http://<user>:<passwd>@<host>/<path>       |  N/A          |
    |  https://<user>:<passwd>@<host>/<path>      |  N/A          |
    |  ftp://<user>:<passwd>@<host>/<path>        |  N/A          |
    |  sftp://<user>:<passwd>@<host>/<path>       |  N/A          |
    |  scp://<user>:<passwd>@<host>/<path>        |  N/A          |
    |  tftp://<host>/<path>                       |  N/A          |
    |  git+https://<user>:<passwd>@<host>/<path>  |  N/A          |
`,
		},

		"source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source IP address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
    |  ipv6    |  IPv6 source address  |
`,
			Description: `Source IP address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
    |  ipv6    |  IPv6 source address  |
`,
		},
	}
}
