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
var _ helpers.VyosTopResourceDataModel = &ProtocolsMplsLdpExportIPvsixExportFilter{}

// ProtocolsMplsLdpExportIPvsixExportFilter describes the resource data model.
type ProtocolsMplsLdpExportIPvsixExportFilter struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsMplsLdpExportIPvsixExportFilterFilterAccessListsix   types.Number `tfsdk:"filter_access_list6" vyos:"filter-access-list6,omitempty"`
	LeafProtocolsMplsLdpExportIPvsixExportFilterNeighborAccessListsix types.Number `tfsdk:"neighbor_access_list6" vyos:"neighbor-access-list6,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsMplsLdpExportIPvsixExportFilter) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsMplsLdpExportIPvsixExportFilter) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsMplsLdpExportIPvsixExportFilter) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpExportIPvsixExportFilter) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"export-filter",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsMplsLdpExportIPvsixExportFilter) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"mpls",

		"ldp",

		"export",

		"ipv6",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsMplsLdpExportIPvsixExportFilter) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpExportIPvsixExportFilter) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"filter_access_list6": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list6 number to apply FEC filtering

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-2699  |  Access list number  |
`,
			Description: `Access-list6 number to apply FEC filtering

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-2699  |  Access list number  |
`,
		},

		"neighbor_access_list6": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list6 number for IPv6 neighbor selection to apply filtering

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-2699  |  Access list number  |
`,
			Description: `Access-list6 number for IPv6 neighbor selection to apply filtering

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-2699  |  Access list number  |
`,
		},
	}
}
