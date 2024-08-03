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
var _ helpers.VyosTopResourceDataModel = &VpnSstpPppOptions{}

// VpnSstpPppOptions describes the resource data model.
type VpnSstpPppOptions struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnSstpPppOptionsMinMtu                      types.String `tfsdk:"min_mtu" vyos:"min-mtu,omitempty"`
	LeafVpnSstpPppOptionsMru                         types.String `tfsdk:"mru" vyos:"mru,omitempty"`
	LeafVpnSstpPppOptionsDisableCcp                  types.Bool   `tfsdk:"disable_ccp" vyos:"disable-ccp,omitempty"`
	LeafVpnSstpPppOptionsMppe                        types.String `tfsdk:"mppe" vyos:"mppe,omitempty"`
	LeafVpnSstpPppOptionsLcpEchoInterval             types.String `tfsdk:"lcp_echo_interval" vyos:"lcp-echo-interval,omitempty"`
	LeafVpnSstpPppOptionsLcpEchoFailure              types.String `tfsdk:"lcp_echo_failure" vyos:"lcp-echo-failure,omitempty"`
	LeafVpnSstpPppOptionsLcpEchoTimeout              types.String `tfsdk:"lcp_echo_timeout" vyos:"lcp-echo-timeout,omitempty"`
	LeafVpnSstpPppOptionsInterfaceCache              types.Number `tfsdk:"interface_cache" vyos:"interface-cache,omitempty"`
	LeafVpnSstpPppOptionsIPvfour                     types.String `tfsdk:"ipv4" vyos:"ipv4,omitempty"`
	LeafVpnSstpPppOptionsIPvsix                      types.String `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	LeafVpnSstpPppOptionsIPvsixInterfaceID           types.String `tfsdk:"ipv6_interface_id" vyos:"ipv6-interface-id,omitempty"`
	LeafVpnSstpPppOptionsIPvsixPeerInterfaceID       types.String `tfsdk:"ipv6_peer_interface_id" vyos:"ipv6-peer-interface-id,omitempty"`
	LeafVpnSstpPppOptionsIPvsixAcceptPeerInterfaceID types.Bool   `tfsdk:"ipv6_accept_peer_interface_id" vyos:"ipv6-accept-peer-interface-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *VpnSstpPppOptions) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnSstpPppOptions) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnSstpPppOptions) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnSstpPppOptions) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"ppp-options",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnSstpPppOptions) GetVyosParentPath() []string {
	return []string{
		"vpn",

		"sstp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *VpnSstpPppOptions) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnSstpPppOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"min_mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Minimum acceptable MTU (68-65535)

`,
			Description: `Minimum acceptable MTU (68-65535)

`,
		},

		"mru": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Preferred MRU (68-65535)

`,
			Description: `Preferred MRU (68-65535)

`,
		},

		"disable_ccp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable Compression Control Protocol (CCP)

`,
			Description: `Disable Compression Control Protocol (CCP)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mppe": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies mppe negotiation preferences

    |  Format   |  Description                                                |
    |-----------|-------------------------------------------------------------|
    |  require  |  send mppe request, if client rejects, drop the connection  |
    |  prefer   |  send mppe request, if client rejects continue              |
    |  deny     |  drop all mppe                                              |
`,
			Description: `Specifies mppe negotiation preferences

    |  Format   |  Description                                                |
    |-----------|-------------------------------------------------------------|
    |  require  |  send mppe request, if client rejects, drop the connection  |
    |  prefer   |  send mppe request, if client rejects continue              |
    |  deny     |  drop all mppe                                              |
`,

			// Default:          stringdefault.StaticString(`prefer`),
			Computed: true,
		},

		"lcp_echo_interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `LCP echo-requests/sec

`,
			Description: `LCP echo-requests/sec

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"lcp_echo_failure": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of Echo-Requests may be sent without valid reply

`,
			Description: `Maximum number of Echo-Requests may be sent without valid reply

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"lcp_echo_timeout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Timeout in seconds to wait for any peer activity. If this option specified it turns on adaptive lcp echo functionality and "lcp-echo-failure" is not used.

`,
			Description: `Timeout in seconds to wait for any peer activity. If this option specified it turns on adaptive lcp echo functionality and "lcp-echo-failure" is not used.

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"interface_cache": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `PPP interface cache

    |  Format    |  Description                           |
    |------------|----------------------------------------|
    |  1-256000  |  Count of interfaces to keep in cache  |
`,
			Description: `PPP interface cache

    |  Format    |  Description                           |
    |------------|----------------------------------------|
    |  1-256000  |  Count of interfaces to keep in cache  |
`,
		},

		"ipv4": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 (IPCP) negotiation algorithm

    |  Format   |  Description                                                 |
    |-----------|--------------------------------------------------------------|
    |  deny     |  Do not negotiate IPv4                                       |
    |  allow    |  Negotiate IPv4 only if client requests                      |
    |  prefer   |  Ask client for IPv4 negotiation, do not fail if it rejects  |
    |  require  |  Require IPv4 negotiation                                    |
`,
			Description: `IPv4 (IPCP) negotiation algorithm

    |  Format   |  Description                                                 |
    |-----------|--------------------------------------------------------------|
    |  deny     |  Do not negotiate IPv4                                       |
    |  allow    |  Negotiate IPv4 only if client requests                      |
    |  prefer   |  Ask client for IPv4 negotiation, do not fail if it rejects  |
    |  require  |  Require IPv4 negotiation                                    |
`,
		},

		"ipv6": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 (IPCP6) negotiation algorithm

    |  Format   |  Description                                                 |
    |-----------|--------------------------------------------------------------|
    |  deny     |  Do not negotiate IPv6                                       |
    |  allow    |  Negotiate IPv6 only if client requests                      |
    |  prefer   |  Ask client for IPv6 negotiation, do not fail if it rejects  |
    |  require  |  Require IPv6 negotiation                                    |
`,
			Description: `IPv6 (IPCP6) negotiation algorithm

    |  Format   |  Description                                                 |
    |-----------|--------------------------------------------------------------|
    |  deny     |  Do not negotiate IPv6                                       |
    |  allow    |  Negotiate IPv6 only if client requests                      |
    |  prefer   |  Ask client for IPv6 negotiation, do not fail if it rejects  |
    |  require  |  Require IPv6 negotiation                                    |
`,

			// Default:          stringdefault.StaticString(`deny`),
			Computed: true,
		},

		"ipv6_interface_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Fixed or random interface identifier for IPv6

    |  Format   |  Description                            |
    |-----------|-----------------------------------------|
    |  random   |  Random interface identifier for IPv6   |
    |  x:x:x:x  |  specify interface identifier for IPv6  |
`,
			Description: `Fixed or random interface identifier for IPv6

    |  Format   |  Description                            |
    |-----------|-----------------------------------------|
    |  random   |  Random interface identifier for IPv6   |
    |  x:x:x:x  |  specify interface identifier for IPv6  |
`,
		},

		"ipv6_peer_interface_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer interface identifier for IPv6

    |  Format       |  Description                                                                |
    |---------------|-----------------------------------------------------------------------------|
    |  x:x:x:x      |  Interface identifier for IPv6                                              |
    |  random       |  Use a random interface identifier for IPv6                                 |
    |  ipv4-addr    |  Calculate interface identifier from IPv4 address, for example 192:168:0:1  |
    |  calling-sid  |  Calculate interface identifier from calling-station-id                     |
`,
			Description: `Peer interface identifier for IPv6

    |  Format       |  Description                                                                |
    |---------------|-----------------------------------------------------------------------------|
    |  x:x:x:x      |  Interface identifier for IPv6                                              |
    |  random       |  Use a random interface identifier for IPv6                                 |
    |  ipv4-addr    |  Calculate interface identifier from IPv4 address, for example 192:168:0:1  |
    |  calling-sid  |  Calculate interface identifier from calling-station-id                     |
`,
		},

		"ipv6_accept_peer_interface_id": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Accept peer interface identifier

`,
			Description: `Accept peer interface identifier

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
