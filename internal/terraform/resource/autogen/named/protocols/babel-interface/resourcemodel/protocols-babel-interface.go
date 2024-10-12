// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsBabelInterface{}

// ProtocolsBabelInterface describes the resource data model.
type ProtocolsBabelInterface struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBabelInterfaceType             types.String `tfsdk:"type" vyos:"type,omitempty"`
	LeafProtocolsBabelInterfaceSplitHorizon     types.String `tfsdk:"split_horizon" vyos:"split-horizon,omitempty"`
	LeafProtocolsBabelInterfaceHelloInterval    types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafProtocolsBabelInterfaceUpdateInterval   types.Number `tfsdk:"update_interval" vyos:"update-interval,omitempty"`
	LeafProtocolsBabelInterfaceRxcost           types.Number `tfsdk:"rxcost" vyos:"rxcost,omitempty"`
	LeafProtocolsBabelInterfaceRttDecay         types.Number `tfsdk:"rtt_decay" vyos:"rtt-decay,omitempty"`
	LeafProtocolsBabelInterfaceRttMin           types.Number `tfsdk:"rtt_min" vyos:"rtt-min,omitempty"`
	LeafProtocolsBabelInterfaceRttMax           types.Number `tfsdk:"rtt_max" vyos:"rtt-max,omitempty"`
	LeafProtocolsBabelInterfaceMaxRttPenalty    types.Number `tfsdk:"max_rtt_penalty" vyos:"max-rtt-penalty,omitempty"`
	LeafProtocolsBabelInterfaceEnableTimestamps types.Bool   `tfsdk:"enable_timestamps" vyos:"enable-timestamps,omitempty"`
	LeafProtocolsBabelInterfaceChannel          types.String `tfsdk:"channel" vyos:"channel,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsBabelInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBabelInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBabelInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBabelInterface) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"babel",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBabelInterface) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"interface": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Interface name

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
					Description: `Interface name

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in interface, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  interface, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface type

    |  Format    |  Description                          |
    |------------|---------------------------------------|
    |  auto      |  Automatically detect interface type  |
    |  wired     |  Wired interface                      |
    |  wireless  |  Wireless interface                   |
`,
			Description: `Interface type

    |  Format    |  Description                          |
    |------------|---------------------------------------|
    |  auto      |  Automatically detect interface type  |
    |  wired     |  Wired interface                      |
    |  wireless  |  Wireless interface                   |
`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		"split_horizon": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Split horizon parameters

    |  Format   |  Description                                                     |
    |-----------|------------------------------------------------------------------|
    |  default  |  Enable on wired interfaces, and disable on wireless interfaces  |
    |  enable   |  Enable split horizon processing                                 |
    |  disable  |  Disable split horizon processing                                |
`,
			Description: `Split horizon parameters

    |  Format   |  Description                                                     |
    |-----------|------------------------------------------------------------------|
    |  default  |  Enable on wired interfaces, and disable on wireless interfaces  |
    |  enable   |  Enable split horizon processing                                 |
    |  disable  |  Disable split horizon processing                                |
`,

			// Default:          stringdefault.StaticString(`default`),
			Computed: true,
		},

		"hello_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time between scheduled hellos

    |  Format     |  Description   |
    |-------------|----------------|
    |  20-655340  |  Milliseconds  |
`,
			Description: `Time between scheduled hellos

    |  Format     |  Description   |
    |-------------|----------------|
    |  20-655340  |  Milliseconds  |
`,

			// Default:          stringdefault.StaticString(`4000`),
			Computed: true,
		},

		"update_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time between scheduled updates

    |  Format     |  Description   |
    |-------------|----------------|
    |  20-655340  |  Milliseconds  |
`,
			Description: `Time between scheduled updates

    |  Format     |  Description   |
    |-------------|----------------|
    |  20-655340  |  Milliseconds  |
`,

			// Default:          stringdefault.StaticString(`20000`),
			Computed: true,
		},

		"rxcost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Base receive cost for this interface

    |  Format   |  Description        |
    |-----------|---------------------|
    |  1-65534  |  Base receive cost  |
`,
			Description: `Base receive cost for this interface

    |  Format   |  Description        |
    |-----------|---------------------|
    |  1-65534  |  Base receive cost  |
`,
		},

		"rtt_decay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Decay factor for exponential moving average of RTT samples

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  1-256   |  Decay factor, in units of 1/256  |
`,
			Description: `Decay factor for exponential moving average of RTT samples

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  1-256   |  Decay factor, in units of 1/256  |
`,

			// Default:          stringdefault.StaticString(`42`),
			Computed: true,
		},

		"rtt_min": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum RTT

    |  Format   |  Description   |
    |-----------|----------------|
    |  1-65535  |  Milliseconds  |
`,
			Description: `Minimum RTT

    |  Format   |  Description   |
    |-----------|----------------|
    |  1-65535  |  Milliseconds  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"rtt_max": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum RTT

    |  Format   |  Description   |
    |-----------|----------------|
    |  1-65535  |  Milliseconds  |
`,
			Description: `Maximum RTT

    |  Format   |  Description   |
    |-----------|----------------|
    |  1-65535  |  Milliseconds  |
`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"max_rtt_penalty": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum additional cost due to RTT

    |  Format   |  Description                                            |
    |-----------|---------------------------------------------------------|
    |  0-65535  |  Milliseconds (0 to disable the use of RTT-based cost)  |
`,
			Description: `Maximum additional cost due to RTT

    |  Format   |  Description                                            |
    |-----------|---------------------------------------------------------|
    |  0-65535  |  Milliseconds (0 to disable the use of RTT-based cost)  |
`,

			// Default:          stringdefault.StaticString(`150`),
			Computed: true,
		},

		"enable_timestamps": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable timestamps with each Hello and IHU message in order to compute RTT values

`,
			Description: `Enable timestamps with each Hello and IHU message in order to compute RTT values

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"channel": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Channel number for diversity routing

    |  Format           |  Description                                                                                                         |
    |-------------------|----------------------------------------------------------------------------------------------------------------------|
    |  1-254            |  Interfaces with a channel number interfere with interfering interfaces and interfaces with the same channel number  |
    |  interfering      |  Interfering interfaces are assumed to interfere with all other channels except non-interfering channels             |
    |  non-interfering  |  Non-interfering interfaces only interfere with themselves                                                           |
`,
			Description: `Channel number for diversity routing

    |  Format           |  Description                                                                                                         |
    |-------------------|----------------------------------------------------------------------------------------------------------------------|
    |  1-254            |  Interfaces with a channel number interfere with interfering interfaces and interfaces with the same channel number  |
    |  interfering      |  Interfering interfaces are assumed to interfere with all other channels except non-interfering channels             |
    |  non-interfering  |  Non-interfering interfaces only interfere with themselves                                                           |
`,
		},

		// Nodes

	}
}
