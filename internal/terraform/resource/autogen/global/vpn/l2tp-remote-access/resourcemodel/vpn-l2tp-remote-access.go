/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (remote-access) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VpnLtwotpRemoteAccess{}

// VpnLtwotpRemoteAccess describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type VpnLtwotpRemoteAccess struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnLtwotpRemoteAccessOutsIDeAddress        types.String `tfsdk:"outside_address" vyos:"outside-address,omitempty"`
	LeafVpnLtwotpRemoteAccessDefaultPool           types.String `tfsdk:"default_pool" vyos:"default-pool,omitempty"`
	LeafVpnLtwotpRemoteAccessDefaultIPvsixPool     types.String `tfsdk:"default_ipv6_pool" vyos:"default-ipv6-pool,omitempty"`
	LeafVpnLtwotpRemoteAccessGatewayAddress        types.String `tfsdk:"gateway_address" vyos:"gateway-address,omitempty"`
	LeafVpnLtwotpRemoteAccessMaxConcurrentSessions types.Number `tfsdk:"max_concurrent_sessions" vyos:"max-concurrent-sessions,omitempty"`
	LeafVpnLtwotpRemoteAccessMtu                   types.String `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafVpnLtwotpRemoteAccessWinsServer            types.List   `tfsdk:"wins_server" vyos:"wins-server,omitempty"`
	LeafVpnLtwotpRemoteAccessDescrIPtion           types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafVpnLtwotpRemoteAccessNameServer            types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`

	// TagNodes

	ExistsTagVpnLtwotpRemoteAccessClientIPPool bool `tfsdk:"-" vyos:"client-ip-pool,child"`

	ExistsTagVpnLtwotpRemoteAccessClientIPvsixPool bool `tfsdk:"-" vyos:"client-ipv6-pool,child"`

	// Nodes

	ExistsNodeVpnLtwotpRemoteAccessAuthentication bool `tfsdk:"-" vyos:"authentication,child"`

	ExistsNodeVpnLtwotpRemoteAccessIPsecSettings bool `tfsdk:"-" vyos:"ipsec-settings,child"`

	ExistsNodeVpnLtwotpRemoteAccessLns bool `tfsdk:"-" vyos:"lns,child"`

	ExistsNodeVpnLtwotpRemoteAccessExtendedScrIPts bool `tfsdk:"-" vyos:"extended-scripts,child"`

	ExistsNodeVpnLtwotpRemoteAccessLimits bool `tfsdk:"-" vyos:"limits,child"`

	ExistsNodeVpnLtwotpRemoteAccessPppOptions bool `tfsdk:"-" vyos:"ppp-options,child"`

	ExistsNodeVpnLtwotpRemoteAccessShaper bool `tfsdk:"-" vyos:"shaper,child"`

	ExistsNodeVpnLtwotpRemoteAccessSnmp bool `tfsdk:"-" vyos:"snmp,child"`

	ExistsNodeVpnLtwotpRemoteAccessLog bool `tfsdk:"-" vyos:"log,child"`
}

// SetID configures the resource ID
func (o *VpnLtwotpRemoteAccess) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnLtwotpRemoteAccess) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnLtwotpRemoteAccess) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnLtwotpRemoteAccess) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"remote-access",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnLtwotpRemoteAccess) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (l2tp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vpn) */
		"vpn", // Node

		"l2tp", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnLtwotpRemoteAccess) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (l2tp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (vpn) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnLtwotpRemoteAccess) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"outside_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (outside-address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `External IP address to which VPN clients will connect

`,
			Description: `External IP address to which VPN clients will connect

`,
		},

		"default_pool":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-pool) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default client IP pool name

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Default IP pool  |
`,
			Description: `Default client IP pool name

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Default IP pool  |
`,
		},

		"default_ipv6_pool":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-ipv6-pool) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default client IPv6 pool name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  Default IPv6 pool  |
`,
			Description: `Default client IPv6 pool name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  Default IPv6 pool  |
`,
		},

		"gateway_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (gateway-address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway IP address

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  ipv4    |  Default Gateway send to the client  |
`,
			Description: `Gateway IP address

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  ipv4    |  Default Gateway send to the client  |
`,
		},

		"max_concurrent_sessions":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (max-concurrent-sessions) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of concurrent session start attempts

    |  Format   |  Description                                          |
    |-----------|-------------------------------------------------------|
    |  0-65535  |  Maximum number of concurrent session start attempts  |
`,
			Description: `Maximum number of concurrent session start attempts

    |  Format   |  Description                                          |
    |-----------|-------------------------------------------------------|
    |  0-65535  |  Maximum number of concurrent session start attempts  |
`,
		},

		"mtu":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mtu) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

`,
			Description: `Maximum Transmission Unit (MTU)

`,

			// Default:          stringdefault.StaticString(`1436`),
			Computed: true,
		},

		"wins_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (wins-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Windows Internet Name Service (WINS) servers propagated to client

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
`,
			Description: `Windows Internet Name Service (WINS) servers propagated to client

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
`,
		},

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (description) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"name_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (name-server) */
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

		// TagNodes

		// Nodes

	}
}
