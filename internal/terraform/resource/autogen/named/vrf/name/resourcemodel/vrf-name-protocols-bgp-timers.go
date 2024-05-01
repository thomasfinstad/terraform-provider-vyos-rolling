// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpTimers{}

// VrfNameProtocolsBgpTimers describes the resource data model.
type VrfNameProtocolsBgpTimers struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpTimersHoldtime  types.String `tfsdk:"holdtime" vyos:"holdtime,omitempty"`
	LeafVrfNameProtocolsBgpTimersKeepalive types.Number `tfsdk:"keepalive" vyos:"keepalive,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpTimers) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"holdtime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hold timer

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  1-65535  |  Hold timer in seconds  |
    |  0        |  Disable hold timer     |
`,
			Description: `Hold timer

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  1-65535  |  Hold timer in seconds  |
    |  0        |  Disable hold timer     |
`,
		},

		"keepalive": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP keepalive interval for this neighbor

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Keepalive interval in seconds  |
`,
			Description: `BGP keepalive interval for this neighbor

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Keepalive interval in seconds  |
`,
		},

		// Nodes

	}
}
