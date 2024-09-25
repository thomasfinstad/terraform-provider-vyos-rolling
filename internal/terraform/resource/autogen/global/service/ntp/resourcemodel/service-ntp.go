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
var _ helpers.VyosTopResourceDataModel = &ServiceNtp{}

// ServiceNtp describes the resource data model.
type ServiceNtp struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceNtpInterface     types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafServiceNtpListenAddress types.List   `tfsdk:"listen_address" vyos:"listen-address,omitempty"`
	LeafServiceNtpVrf           types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafServiceNtpLeapSecond    types.String `tfsdk:"leap_second" vyos:"leap-second,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceNtpServer bool `tfsdk:"-" vyos:"server,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceNtpAllowClient bool `tfsdk:"-" vyos:"allow-client,child"`
	ExistsNodeServiceNtpPtp         bool `tfsdk:"-" vyos:"ptp,child"`
}

// SetID configures the resource ID
func (o *ServiceNtp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceNtp) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceNtp) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceNtp) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"ntp",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceNtp) GetVyosParentPath() []string {
	return []string{
		"service",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ServiceNtp) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceNtp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"interface": schema.StringAttribute{
			Optional: true,
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

		"listen_address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IP addresses to listen on

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming connections  |
    |  ipv6    |  IPv6 address to listen for incoming connections  |
`,
			Description: `Local IP addresses to listen on

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming connections  |
    |  ipv6    |  IPv6 address to listen for incoming connections  |
`,
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

		"leap_second": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Leap second behavior

    |  Format    |  Description                                                                  |
    |------------|-------------------------------------------------------------------------------|
    |  ignore    |  No correction is applied to the clock for the leap second                    |
    |  smear     |  Correct served time slowly be slewing instead of stepping                    |
    |  system    |  Kernel steps the system clock forward or backward                            |
    |  timezone  |  Use UTC timezone database to determine when will the next leap second occur  |
`,
			Description: `Leap second behavior

    |  Format    |  Description                                                                  |
    |------------|-------------------------------------------------------------------------------|
    |  ignore    |  No correction is applied to the clock for the leap second                    |
    |  smear     |  Correct served time slowly be slewing instead of stepping                    |
    |  system    |  Kernel steps the system clock forward or backward                            |
    |  timezone  |  Use UTC timezone database to determine when will the next leap second occur  |
`,

			// Default:          stringdefault.StaticString(`timezone`),
			Computed: true,
		},
	}
}
