// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBfdProfileInterval{}

// ProtocolsBfdProfileInterval describes the resource data model.
type ProtocolsBfdProfileInterval struct {
	// LeafNodes
	LeafProtocolsBfdProfileIntervalReceive      types.Number `tfsdk:"receive" vyos:"receive,omitempty"`
	LeafProtocolsBfdProfileIntervalTransmit     types.Number `tfsdk:"transmit" vyos:"transmit,omitempty"`
	LeafProtocolsBfdProfileIntervalMultIPlier   types.Number `tfsdk:"multiplier" vyos:"multiplier,omitempty"`
	LeafProtocolsBfdProfileIntervalEchoInterval types.Number `tfsdk:"echo_interval" vyos:"echo-interval,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBfdProfileInterval) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"receive": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval of receiving control packets

    |  Format    |  Description               |
    |------------|----------------------------|
    |  10-60000  |  Interval in milliseconds  |
`,
			Description: `Minimum interval of receiving control packets

    |  Format    |  Description               |
    |------------|----------------------------|
    |  10-60000  |  Interval in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"transmit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval of transmitting control packets

    |  Format    |  Description               |
    |------------|----------------------------|
    |  10-60000  |  Interval in milliseconds  |
`,
			Description: `Minimum interval of transmitting control packets

    |  Format    |  Description               |
    |------------|----------------------------|
    |  10-60000  |  Interval in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"multiplier": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Multiplier to determine packet loss

    |  Format  |  Description                                                    |
    |----------|-----------------------------------------------------------------|
    |  2-255   |  Remote transmission interval will be multiplied by this value  |
`,
			Description: `Multiplier to determine packet loss

    |  Format  |  Description                                                    |
    |----------|-----------------------------------------------------------------|
    |  2-255   |  Remote transmission interval will be multiplied by this value  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"echo_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Echo receive transmission interval

    |  Format    |  Description                                                                             |
    |------------|------------------------------------------------------------------------------------------|
    |  10-60000  |  The minimal echo receive transmission interval that this system is capable of handling  |
`,
			Description: `Echo receive transmission interval

    |  Format    |  Description                                                                             |
    |------------|------------------------------------------------------------------------------------------|
    |  10-60000  |  The minimal echo receive transmission interval that this system is capable of handling  |
`,
		},

		// Nodes

	}
}
