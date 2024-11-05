/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VpnOpenconnectNetworkSettings{}

// VpnOpenconnectNetworkSettings describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type VpnOpenconnectNetworkSettings struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnOpenconnectNetworkSettingsPushRoute    types.List   `tfsdk:"push_route" vyos:"push-route,omitempty"`
	LeafVpnOpenconnectNetworkSettingsNameServer   types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`
	LeafVpnOpenconnectNetworkSettingsSplitDNS     types.List   `tfsdk:"split_dns" vyos:"split-dns,omitempty"`
	LeafVpnOpenconnectNetworkSettingsTunnelAllDNS types.String `tfsdk:"tunnel_all_dns" vyos:"tunnel-all-dns,omitempty"`

	// TagNodes

	// Nodes

	ExistsNodeVpnOpenconnectNetworkSettingsClientIPSettings bool `tfsdk:"-" vyos:"client-ip-settings,child"`

	ExistsNodeVpnOpenconnectNetworkSettingsClientIPvsixPool bool `tfsdk:"-" vyos:"client-ipv6-pool,child"`
}

// SetID configures the resource ID
func (o *VpnOpenconnectNetworkSettings) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnOpenconnectNetworkSettings) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnOpenconnectNetworkSettings) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnOpenconnectNetworkSettings) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"network-settings",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnOpenconnectNetworkSettings) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"vpn", // Node

		"openconnect", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnOpenconnectNetworkSettings) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnOpenconnectNetworkSettings) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"push_route":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Route to be pushed to the client

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 network and prefix length  |
    |  ipv6net  |  IPv6 network and prefix length  |
`,
			Description: `Route to be pushed to the client

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 network and prefix length  |
    |  ipv6net  |  IPv6 network and prefix length  |
`,
		},

		"name_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
    |  ipv6    |  Domain Name Server (DNS) IPv6 address  |
`,
			Description: `Domain Name Servers (DNS) addresses

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
    |  ipv6    |  Domain Name Server (DNS) IPv6 address  |
`,
		},

		"split_dns":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domains over which the provided DNS should be used

    |  Format  |  Description           |
    |----------|------------------------|
    |  txt     |  Client prefix length  |
`,
			Description: `Domains over which the provided DNS should be used

    |  Format  |  Description           |
    |----------|------------------------|
    |  txt     |  Client prefix length  |
`,
		},

		"tunnel_all_dns":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `If the tunnel-all-dns option is set to yes, tunnel all DNS queries via the VPN. This is the default when a default route is set.

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  yes     |  Enable tunneling of all DNS traffic   |
    |  no      |  Disable tunneling of all DNS traffic  |
`,
			Description: `If the tunnel-all-dns option is set to yes, tunnel all DNS queries via the VPN. This is the default when a default route is set.

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  yes     |  Enable tunneling of all DNS traffic   |
    |  no      |  Disable tunneling of all DNS traffic  |
`,

			// Default:          stringdefault.StaticString(`no`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
