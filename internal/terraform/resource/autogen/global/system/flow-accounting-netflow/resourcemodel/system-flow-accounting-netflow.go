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
var _ helpers.VyosTopResourceDataModel = &SystemFlowAccountingNetflow{}

// SystemFlowAccountingNetflow describes the resource data model.
type SystemFlowAccountingNetflow struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemFlowAccountingNetflowEngineID      types.String `tfsdk:"engine_id" vyos:"engine-id,omitempty"`
	LeafSystemFlowAccountingNetflowMaxFlows      types.Number `tfsdk:"max_flows" vyos:"max-flows,omitempty"`
	LeafSystemFlowAccountingNetflowSamplingRate  types.Number `tfsdk:"sampling_rate" vyos:"sampling-rate,omitempty"`
	LeafSystemFlowAccountingNetflowSourceAddress types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafSystemFlowAccountingNetflowVersion       types.String `tfsdk:"version" vyos:"version,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagSystemFlowAccountingNetflowServer bool `tfsdk:"-" vyos:"server,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeSystemFlowAccountingNetflowTimeout bool `tfsdk:"-" vyos:"timeout,child"`
}

// SetID configures the resource ID
func (o *SystemFlowAccountingNetflow) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemFlowAccountingNetflow) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemFlowAccountingNetflow) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemFlowAccountingNetflow) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"netflow",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemFlowAccountingNetflow) GetVyosParentPath() []string {
	return []string{
		"system",

		"flow-accounting",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *SystemFlowAccountingNetflow) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemFlowAccountingNetflow) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"engine_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NetFlow engine-id

    |  Format                |  Description                       |
    |------------------------|------------------------------------|
    |  0-255 or 0-255:0-255  |  NetFlow engine-id for v5          |
    |  u32                   |  NetFlow engine-id for v9 / IPFIX  |
`,
			Description: `NetFlow engine-id

    |  Format                |  Description                       |
    |------------------------|------------------------------------|
    |  0-255 or 0-255:0-255  |  NetFlow engine-id for v5          |
    |  u32                   |  NetFlow engine-id for v9 / IPFIX  |
`,
		},

		"max_flows": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `NetFlow maximum flows

    |  Format  |  Description            |
    |----------|-------------------------|
    |  u32     |  NetFlow maximum flows  |
`,
			Description: `NetFlow maximum flows

    |  Format  |  Description            |
    |----------|-------------------------|
    |  u32     |  NetFlow maximum flows  |
`,
		},

		"sampling_rate": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `NetFlow sampling-rate

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  u32     |  Sampling rate (1 in N packets)  |
`,
			Description: `NetFlow sampling-rate

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  u32     |  Sampling rate (1 in N packets)  |
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

		"version": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NetFlow version to export

    |  Format  |  Description                                        |
    |----------|-----------------------------------------------------|
    |  5       |  NetFlow version 5                                  |
    |  9       |  NetFlow version 9                                  |
    |  10      |  Internet Protocol Flow Information Export (IPFIX)  |
`,
			Description: `NetFlow version to export

    |  Format  |  Description                                        |
    |----------|-----------------------------------------------------|
    |  5       |  NetFlow version 5                                  |
    |  9       |  NetFlow version 9                                  |
    |  10      |  Internet Protocol Flow Information Export (IPFIX)  |
`,

			// Default:          stringdefault.StaticString(`9`),
			Computed: true,
		},
	}
}
