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
var _ helpers.VyosTopResourceDataModel = &ProtocolsPimsix{}

// ProtocolsPimsix describes the resource data model.
type ProtocolsPimsix struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsPimsixJoinPruneInterval    types.Number `tfsdk:"join_prune_interval" vyos:"join-prune-interval,omitempty"`
	LeafProtocolsPimsixKeepAliveTimer       types.Number `tfsdk:"keep_alive_timer" vyos:"keep-alive-timer,omitempty"`
	LeafProtocolsPimsixPackets              types.Number `tfsdk:"packets" vyos:"packets,omitempty"`
	LeafProtocolsPimsixRegisterSuppressTime types.Number `tfsdk:"register_suppress_time" vyos:"register-suppress-time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsPimsixInterface bool `tfsdk:"-" vyos:"interface,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsPimsixRp bool `tfsdk:"-" vyos:"rp,child"`
}

// SetID configures the resource ID
func (o *ProtocolsPimsix) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsPimsix) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsPimsix) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsPimsix) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"pim6",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsPimsix) GetVyosParentPath() []string {
	return []string{
		"protocols",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsPimsix) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsPimsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"join_prune_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Join prune send interval

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-65535  |  Interval in seconds  |
`,
			Description: `Join prune send interval

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-65535  |  Interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"keep_alive_timer": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Keep alive Timer

    |  Format   |  Description                  |
    |-----------|-------------------------------|
    |  1-65535  |  Keep alive Timer in seconds  |
`,
			Description: `Keep alive Timer

    |  Format   |  Description                  |
    |-----------|-------------------------------|
    |  1-65535  |  Keep alive Timer in seconds  |
`,
		},

		"packets": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Packets to process at once

    |  Format  |  Description        |
    |----------|---------------------|
    |  1-255   |  Number of packets  |
`,
			Description: `Packets to process at once

    |  Format  |  Description        |
    |----------|---------------------|
    |  1-255   |  Number of packets  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"register_suppress_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Register suppress timer

    |  Format   |  Description       |
    |-----------|--------------------|
    |  1-65535  |  Timer in seconds  |
`,
			Description: `Register suppress timer

    |  Format   |  Description       |
    |-----------|--------------------|
    |  1-65535  |  Timer in seconds  |
`,
		},
	}
}
