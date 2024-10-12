// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ProtocolsMplsLdpExportIPvsix{}

// ProtocolsMplsLdpExportIPvsix describes the resource data model.
type ProtocolsMplsLdpExportIPvsix struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsMplsLdpExportIPvsixExplicitNull types.Bool `tfsdk:"explicit_null" vyos:"explicit-null,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsMplsLdpExportIPvsixExportFilter bool `tfsdk:"-" vyos:"export-filter,child"`
}

// SetID configures the resource ID
func (o *ProtocolsMplsLdpExportIPvsix) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsMplsLdpExportIPvsix) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsMplsLdpExportIPvsix) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpExportIPvsix) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"ipv6",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsMplsLdpExportIPvsix) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"mpls",

		"ldp",

		"export",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsMplsLdpExportIPvsix) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpExportIPvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"explicit_null": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Explicit-Null Label

`,
			Description: `Explicit-Null Label

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
