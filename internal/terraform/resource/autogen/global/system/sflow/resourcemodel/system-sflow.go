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
var _ helpers.VyosTopResourceDataModel = &SystemSflow{}

// SystemSflow describes the resource data model.
type SystemSflow struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemSflowAgentAddress     types.String `tfsdk:"agent_address" vyos:"agent-address,omitempty"`
	LeafSystemSflowAgentInterface   types.String `tfsdk:"agent_interface" vyos:"agent-interface,omitempty"`
	LeafSystemSflowDropMonitorLimit types.Number `tfsdk:"drop_monitor_limit" vyos:"drop-monitor-limit,omitempty"`
	LeafSystemSflowInterface        types.List   `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafSystemSflowPolling          types.Number `tfsdk:"polling" vyos:"polling,omitempty"`
	LeafSystemSflowSamplingRate     types.Number `tfsdk:"sampling_rate" vyos:"sampling-rate,omitempty"`
	LeafSystemSflowVrf              types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagSystemSflowServer bool `tfsdk:"-" vyos:"server,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemSflow) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemSflow) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemSflow) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSflow) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"sflow",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemSflow) GetVyosParentPath() []string {
	return []string{
		"system",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *SystemSflow) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSflow) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"agent_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `sFlow agent IPv4 or IPv6 address

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  sFlow IPv4 agent address  |
    |  ipv6    |  sFlow IPv6 agent address  |
`,
			Description: `sFlow agent IPv4 or IPv6 address

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  sFlow IPv4 agent address  |
    |  ipv6    |  sFlow IPv6 agent address  |
`,
		},

		"agent_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address associated with this interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `IP address associated with this interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		"drop_monitor_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Export headers of dropped by kernel packets

    |  Format   |  Description                                                               |
    |-----------|----------------------------------------------------------------------------|
    |  1-65535  |  Maximum rate limit of N drops per second send out in the sFlow datagrams  |
`,
			Description: `Export headers of dropped by kernel packets

    |  Format   |  Description                                                               |
    |-----------|----------------------------------------------------------------------------|
    |  1-65535  |  Maximum rate limit of N drops per second send out in the sFlow datagrams  |
`,
		},

		"interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		"polling": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Schedule counter-polling in seconds

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-600   |  Polling rate in seconds  |
`,
			Description: `Schedule counter-polling in seconds

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-600   |  Polling rate in seconds  |
`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"sampling_rate": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `sFlow sampling-rate

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  1-65535  |  Sampling rate (1 in N packets)  |
`,
			Description: `sFlow sampling-rate

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  1-65535  |  Sampling rate (1 in N packets)  |
`,

			// Default:          stringdefault.StaticString(`1000`),
			Computed: true,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
		},
	}
}
