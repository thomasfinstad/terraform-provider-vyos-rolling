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
var _ helpers.VyosTopResourceDataModel = &VpnIPsecRemoteAccessRadius{}

// VpnIPsecRemoteAccessRadius describes the resource data model.
type VpnIPsecRemoteAccessRadius struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnIPsecRemoteAccessRadiusSourceAddress types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafVpnIPsecRemoteAccessRadiusNasIDentifier types.String `tfsdk:"nas_identifier" vyos:"nas-identifier,omitempty"`
	LeafVpnIPsecRemoteAccessRadiusTimeout       types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVpnIPsecRemoteAccessRadiusServer bool `tfsdk:"-" vyos:"server,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *VpnIPsecRemoteAccessRadius) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnIPsecRemoteAccessRadius) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnIPsecRemoteAccessRadius) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecRemoteAccessRadius) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"radius",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnIPsecRemoteAccessRadius) GetVyosParentPath() []string {
	return []string{
		"vpn",

		"ipsec",

		"remote-access",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *VpnIPsecRemoteAccessRadius) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecRemoteAccessRadius) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

		"nas_identifier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAS-Identifier attribute sent to RADIUS

`,
			Description: `NAS-Identifier attribute sent to RADIUS

`,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session timeout

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  1-240   |  Session timeout in seconds (default: 2)  |
`,
			Description: `Session timeout

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  1-240   |  Session timeout in seconds (default: 2)  |
`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},
	}
}
