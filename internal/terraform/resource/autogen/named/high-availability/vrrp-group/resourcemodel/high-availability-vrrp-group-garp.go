// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &HighAvailabilityVrrpGroupGarp{}

// HighAvailabilityVrrpGroupGarp describes the resource data model.
type HighAvailabilityVrrpGroupGarp struct {
	// LeafNodes
	LeafHighAvailabilityVrrpGroupGarpInterval            types.String `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafHighAvailabilityVrrpGroupGarpMasterDelay         types.Number `tfsdk:"master_delay" vyos:"master-delay,omitempty"`
	LeafHighAvailabilityVrrpGroupGarpMasterRefresh       types.Number `tfsdk:"master_refresh" vyos:"master-refresh,omitempty"`
	LeafHighAvailabilityVrrpGroupGarpMasterRefreshRepeat types.Number `tfsdk:"master_refresh_repeat" vyos:"master-refresh-repeat,omitempty"`
	LeafHighAvailabilityVrrpGroupGarpMasterRepeat        types.Number `tfsdk:"master_repeat" vyos:"master-repeat,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroupGarp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval between Gratuitous ARP

    |  Format        |  Description                                   |
    |----------------|------------------------------------------------|
    |  <0.000-1000>  |  Interval in seconds, resolution microseconds  |
`,
			Description: `Interval between Gratuitous ARP

    |  Format        |  Description                                   |
    |----------------|------------------------------------------------|
    |  <0.000-1000>  |  Interval in seconds, resolution microseconds  |
`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"master_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay for second set of gratuitous ARPs after transition to master

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-1000  |  Delay in seconds  |
`,
			Description: `Delay for second set of gratuitous ARPs after transition to master

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-1000  |  Delay in seconds  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"master_refresh": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum time interval for refreshing gratuitous ARPs while beeing master

    |  Format  |  Description          |
    |----------|-----------------------|
    |  0       |  No refresh           |
    |  1-255   |  Interval in seconds  |
`,
			Description: `Minimum time interval for refreshing gratuitous ARPs while beeing master

    |  Format  |  Description          |
    |----------|-----------------------|
    |  0       |  No refresh           |
    |  1-255   |  Interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"master_refresh_repeat": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of gratuitous ARP messages to send at a time while beeing master

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  1-255   |  Number of gratuitous ARP messages  |
`,
			Description: `Number of gratuitous ARP messages to send at a time while beeing master

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  1-255   |  Number of gratuitous ARP messages  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"master_repeat": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of gratuitous ARP messages to send at a time after transition to master

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  1-255   |  Number of gratuitous ARP messages  |
`,
			Description: `Number of gratuitous ARP messages to send at a time after transition to master

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  1-255   |  Number of gratuitous ARP messages  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// Nodes

	}
}
