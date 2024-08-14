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
var _ helpers.VyosTopResourceDataModel = &VpnOpenconnectNetworkSettingsClientIPvsixPool{}

// VpnOpenconnectNetworkSettingsClientIPvsixPool describes the resource data model.
type VpnOpenconnectNetworkSettingsClientIPvsixPool struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnOpenconnectNetworkSettingsClientIPvsixPoolPrefix types.String `tfsdk:"prefix" vyos:"prefix,omitempty"`
	LeafVpnOpenconnectNetworkSettingsClientIPvsixPoolMask   types.Number `tfsdk:"mask" vyos:"mask,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"client-ipv6-pool",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) GetVyosParentPath() []string {
	return []string{
		"vpn",

		"openconnect",

		"network-settings",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnOpenconnectNetworkSettingsClientIPvsixPool) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pool of addresses used to assign to clients

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `Pool of addresses used to assign to clients

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
		},

		"mask": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length used for individual client

    |  Format  |  Description           |
    |----------|------------------------|
    |  48-128  |  Client prefix length  |
`,
			Description: `Prefix length used for individual client

    |  Format  |  Description           |
    |----------|------------------------|
    |  48-128  |  Client prefix length  |
`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},
	}
}