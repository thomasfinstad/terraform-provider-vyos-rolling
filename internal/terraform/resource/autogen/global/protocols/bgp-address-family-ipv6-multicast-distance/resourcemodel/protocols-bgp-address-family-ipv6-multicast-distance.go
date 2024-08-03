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
var _ helpers.VyosTopResourceDataModel = &ProtocolsBgpAddressFamilyIPvsixMulticastDistance{}

// ProtocolsBgpAddressFamilyIPvsixMulticastDistance describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixMulticastDistance struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixMulticastDistanceExternal types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixMulticastDistanceInternal types.Number `tfsdk:"internal" vyos:"internal,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixMulticastDistanceLocal    types.Number `tfsdk:"local" vyos:"local,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix bool `tfsdk:"-" vyos:"prefix,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastDistance) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastDistance) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastDistance) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastDistance) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"distance",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastDistance) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"bgp",

		"address-family",

		"ipv6-multicast",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastDistance) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixMulticastDistance) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"external": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `eBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  eBGP routes administrative distance  |
`,
			Description: `eBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  eBGP routes administrative distance  |
`,
		},

		"internal": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `iBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  iBGP routes administrative distance  |
`,
			Description: `iBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  iBGP routes administrative distance  |
`,
		},

		"local": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Locally originated BGP routes administrative distance

    |  Format  |  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   |  Locally originated BGP routes administrative distance  |
`,
			Description: `Locally originated BGP routes administrative distance

    |  Format  |  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   |  Locally originated BGP routes administrative distance  |
`,
		},
	}
}
