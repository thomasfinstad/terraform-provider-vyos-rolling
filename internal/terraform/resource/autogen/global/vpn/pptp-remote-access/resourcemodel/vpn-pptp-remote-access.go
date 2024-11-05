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

var _ helpers.VyosTopResourceDataModel = &VpnPptpRemoteAccess{}

// VpnPptpRemoteAccess describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type VpnPptpRemoteAccess struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnPptpRemoteAccessOutsIDeAddress        types.String `tfsdk:"outside_address" vyos:"outside-address,omitempty"`
	LeafVpnPptpRemoteAccessDefaultPool           types.String `tfsdk:"default_pool" vyos:"default-pool,omitempty"`
	LeafVpnPptpRemoteAccessDefaultIPvsixPool     types.String `tfsdk:"default_ipv6_pool" vyos:"default-ipv6-pool,omitempty"`
	LeafVpnPptpRemoteAccessGatewayAddress        types.String `tfsdk:"gateway_address" vyos:"gateway-address,omitempty"`
	LeafVpnPptpRemoteAccessMaxConcurrentSessions types.Number `tfsdk:"max_concurrent_sessions" vyos:"max-concurrent-sessions,omitempty"`
	LeafVpnPptpRemoteAccessMtu                   types.String `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafVpnPptpRemoteAccessWinsServer            types.List   `tfsdk:"wins_server" vyos:"wins-server,omitempty"`
	LeafVpnPptpRemoteAccessDescrIPtion           types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafVpnPptpRemoteAccessNameServer            types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`

	// TagNodes

	ExistsTagVpnPptpRemoteAccessClientIPPool bool `tfsdk:"-" vyos:"client-ip-pool,child"`

	ExistsTagVpnPptpRemoteAccessClientIPvsixPool bool `tfsdk:"-" vyos:"client-ipv6-pool,child"`

	// Nodes

	ExistsNodeVpnPptpRemoteAccessAuthentication bool `tfsdk:"-" vyos:"authentication,child"`

	ExistsNodeVpnPptpRemoteAccessExtendedScrIPts bool `tfsdk:"-" vyos:"extended-scripts,child"`

	ExistsNodeVpnPptpRemoteAccessLimits bool `tfsdk:"-" vyos:"limits,child"`

	ExistsNodeVpnPptpRemoteAccessPppOptions bool `tfsdk:"-" vyos:"ppp-options,child"`

	ExistsNodeVpnPptpRemoteAccessShaper bool `tfsdk:"-" vyos:"shaper,child"`

	ExistsNodeVpnPptpRemoteAccessSnmp bool `tfsdk:"-" vyos:"snmp,child"`

	ExistsNodeVpnPptpRemoteAccessLog bool `tfsdk:"-" vyos:"log,child"`
}

// SetID configures the resource ID
func (o *VpnPptpRemoteAccess) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnPptpRemoteAccess) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnPptpRemoteAccess) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnPptpRemoteAccess) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"remote-access",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnPptpRemoteAccess) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"vpn", // Node

		"pptp", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnPptpRemoteAccess) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnPptpRemoteAccess) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `External IP address to which VPN clients will connect

`,
			Description: `External IP address to which VPN clients will connect

`,
		},

		"default_pool":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
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
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
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

		// TagNodes

		// Nodes

	}
}
