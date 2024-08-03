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
var _ helpers.VyosTopResourceDataModel = &ServiceIPoeServerAuthenticationRadius{}

// ServiceIPoeServerAuthenticationRadius describes the resource data model.
type ServiceIPoeServerAuthenticationRadius struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceIPoeServerAuthenticationRadiusSourceAddress             types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusAccountingInterimInterval types.Number `tfsdk:"accounting_interim_interval" vyos:"accounting-interim-interval,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusAcctInterimJitter         types.Number `tfsdk:"acct_interim_jitter" vyos:"acct-interim-jitter,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusTimeout                   types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusAcctTimeout               types.Number `tfsdk:"acct_timeout" vyos:"acct-timeout,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusMaxTry                    types.Number `tfsdk:"max_try" vyos:"max-try,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusNasIDentifier             types.String `tfsdk:"nas_identifier" vyos:"nas-identifier,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusNasIPAddress              types.String `tfsdk:"nas_ip_address" vyos:"nas-ip-address,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusPreallocateVif            types.Bool   `tfsdk:"preallocate_vif" vyos:"preallocate-vif,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceIPoeServerAuthenticationRadiusServer bool `tfsdk:"-" vyos:"server,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceIPoeServerAuthenticationRadiusDynamicAuthor bool `tfsdk:"-" vyos:"dynamic-author,child"`
	ExistsNodeServiceIPoeServerAuthenticationRadiusRateLimit     bool `tfsdk:"-" vyos:"rate-limit,child"`
}

// SetID configures the resource ID
func (o *ServiceIPoeServerAuthenticationRadius) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceIPoeServerAuthenticationRadius) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceIPoeServerAuthenticationRadius) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerAuthenticationRadius) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"radius",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceIPoeServerAuthenticationRadius) GetVyosParentPath() []string {
	return []string{
		"service",

		"ipoe-server",

		"authentication",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ServiceIPoeServerAuthenticationRadius) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerAuthenticationRadius) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 source address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
`,
			Description: `IPv4 source address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
`,
		},

		"accounting_interim_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval in seconds to send accounting information

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  1-3600  |  Interval in seconds to send accounting information  |
`,
			Description: `Interval in seconds to send accounting information

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  1-3600  |  Interval in seconds to send accounting information  |
`,
		},

		"acct_interim_jitter": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum jitter value in seconds to be applied to accounting information interval

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  1-60    |  Maximum jitter value in seconds  |
`,
			Description: `Maximum jitter value in seconds to be applied to accounting information interval

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  1-60    |  Maximum jitter value in seconds  |
`,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout in seconds to wait response from RADIUS server

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-60    |  Timeout in seconds  |
`,
			Description: `Timeout in seconds to wait response from RADIUS server

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-60    |  Timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"acct_timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for Interim-Update packets, terminate session afterwards

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  0-60    |  Timeout in seconds, 0 to keep active  |
`,
			Description: `Timeout for Interim-Update packets, terminate session afterwards

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  0-60    |  Timeout in seconds, 0 to keep active  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"max_try": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of tries to send Access-Request/Accounting-Request queries

    |  Format  |  Description    |
    |----------|-----------------|
    |  1-20    |  Maximum tries  |
`,
			Description: `Number of tries to send Access-Request/Accounting-Request queries

    |  Format  |  Description    |
    |----------|-----------------|
    |  1-20    |  Maximum tries  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"nas_identifier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAS-Identifier attribute sent to RADIUS

`,
			Description: `NAS-Identifier attribute sent to RADIUS

`,
		},

		"nas_ip_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAS-IP-Address attribute sent to RADIUS

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  NAS-IP-Address attribute  |
`,
			Description: `NAS-IP-Address attribute sent to RADIUS

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  NAS-IP-Address attribute  |
`,
		},

		"preallocate_vif": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable attribute NAS-Port-Id in Access-Request

`,
			Description: `Enable attribute NAS-Port-Id in Access-Request

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
