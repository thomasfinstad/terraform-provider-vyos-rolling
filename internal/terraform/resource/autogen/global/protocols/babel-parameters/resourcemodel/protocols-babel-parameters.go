// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
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
var _ helpers.VyosTopResourceDataModel = &ProtocolsBabelParameters{}

// ProtocolsBabelParameters describes the resource data model.
type ProtocolsBabelParameters struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBabelParametersDiversity         types.Bool   `tfsdk:"diversity" vyos:"diversity,omitempty"`
	LeafProtocolsBabelParametersDiversityFactor   types.Number `tfsdk:"diversity_factor" vyos:"diversity-factor,omitempty"`
	LeafProtocolsBabelParametersResendDelay       types.Number `tfsdk:"resend_delay" vyos:"resend-delay,omitempty"`
	LeafProtocolsBabelParametersSmoothingHalfLife types.Number `tfsdk:"smoothing_half_life" vyos:"smoothing-half-life,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsBabelParameters) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBabelParameters) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBabelParameters) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelParameters) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"parameters",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBabelParameters) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"babel",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsBabelParameters) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelParameters) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"diversity": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable diversity-aware routing

`,
			Description: `Enable diversity-aware routing

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"diversity_factor": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Multiplicative factor used for diversity routing

    |  Format  |  Description                               |
    |----------|--------------------------------------------|
    |  1-256   |  Multiplicative factor, in units of 1/256  |
`,
			Description: `Multiplicative factor used for diversity routing

    |  Format  |  Description                               |
    |----------|--------------------------------------------|
    |  1-256   |  Multiplicative factor, in units of 1/256  |
`,

			// Default:          stringdefault.StaticString(`256`),
			Computed: true,
		},

		"resend_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time before resending a message

    |  Format     |  Description   |
    |-------------|----------------|
    |  20-655340  |  Milliseconds  |
`,
			Description: `Time before resending a message

    |  Format     |  Description   |
    |-------------|----------------|
    |  20-655340  |  Milliseconds  |
`,

			// Default:          stringdefault.StaticString(`2000`),
			Computed: true,
		},

		"smoothing_half_life": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Smoothing half-life

    |  Format   |  Description  |
    |-----------|---------------|
    |  0-65534  |  Seconds      |
`,
			Description: `Smoothing half-life

    |  Format   |  Description  |
    |-----------|---------------|
    |  0-65534  |  Seconds      |
`,

			// Default:          stringdefault.StaticString(`4`),
			Computed: true,
		},
	}
}
