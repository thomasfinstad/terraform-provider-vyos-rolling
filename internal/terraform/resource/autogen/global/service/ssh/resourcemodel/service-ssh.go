// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
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
var _ helpers.VyosTopResourceDataModel = &ServiceTCP{}

// ServiceTCP describes the resource data model.
type ServiceTCP struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceTCPCIPhers                       types.List   `tfsdk:"ciphers" vyos:"ciphers,omitempty"`
	LeafServiceTCPDisableHostValIDation         types.Bool   `tfsdk:"disable_host_validation" vyos:"disable-host-validation,omitempty"`
	LeafServiceTCPDisablePasswordAuthentication types.Bool   `tfsdk:"disable_password_authentication" vyos:"disable-password-authentication,omitempty"`
	LeafServiceTCPHostkeyAlgorithm              types.List   `tfsdk:"hostkey_algorithm" vyos:"hostkey-algorithm,omitempty"`
	LeafServiceTCPPubkeyAcceptedAlgorithm       types.List   `tfsdk:"pubkey_accepted_algorithm" vyos:"pubkey-accepted-algorithm,omitempty"`
	LeafServiceTCPKeyExchange                   types.List   `tfsdk:"key_exchange" vyos:"key-exchange,omitempty"`
	LeafServiceTCPListenAddress                 types.List   `tfsdk:"listen_address" vyos:"listen-address,omitempty"`
	LeafServiceTCPLoglevel                      types.String `tfsdk:"loglevel" vyos:"loglevel,omitempty"`
	LeafServiceTCPMac                           types.List   `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafServiceTCPPort                          types.List   `tfsdk:"port" vyos:"port,omitempty"`
	LeafServiceTCPClientKeepaliveInterval       types.Number `tfsdk:"client_keepalive_interval" vyos:"client-keepalive-interval,omitempty"`
	LeafServiceTCPVrf                           types.List   `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceTCPAccessControl     bool `tfsdk:"-" vyos:"access-control,child"`
	ExistsNodeServiceTCPDynamicProtection bool `tfsdk:"-" vyos:"dynamic-protection,child"`
	ExistsNodeServiceTCPRekey             bool `tfsdk:"-" vyos:"rekey,child"`
}

// SetID configures the resource ID
func (o *ServiceTCP) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceTCP) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceTCP) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceTCP) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"ssh",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceTCP) GetVyosParentPath() []string {
	return []string{
		"service",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ServiceTCP) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"ciphers": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Allowed ciphers

`,
			Description: `Allowed ciphers

`,
		},

		"disable_host_validation": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP Address to Hostname lookup

`,
			Description: `Disable IP Address to Hostname lookup

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_password_authentication": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable password-based authentication

`,
			Description: `Disable password-based authentication

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hostkey_algorithm": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Allowed host key signature algorithms

`,
			Description: `Allowed host key signature algorithms

`,
		},

		"pubkey_accepted_algorithm": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Allowed pubkey signature algorithms

`,
			Description: `Allowed pubkey signature algorithms

`,
		},

		"key_exchange": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Allowed key exchange (KEX) algorithms

`,
			Description: `Allowed key exchange (KEX) algorithms

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

		"loglevel": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Log level

    |  Format   |  Description                              |
    |-----------|-------------------------------------------|
    |  quiet    |  stay silent                              |
    |  fatal    |  log fatals only                          |
    |  error    |  log errors and fatals only               |
    |  info     |  default log level                        |
    |  verbose  |  enable logging of failed login attempts  |
`,
			Description: `Log level

    |  Format   |  Description                              |
    |-----------|-------------------------------------------|
    |  quiet    |  stay silent                              |
    |  fatal    |  log fatals only                          |
    |  error    |  log errors and fatals only               |
    |  info     |  default log level                        |
    |  verbose  |  enable logging of failed login attempts  |
`,

			// Default:          stringdefault.StaticString(`info`),
			Computed: true,
		},

		"mac": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Allowed message authentication code (MAC) algorithms

`,
			Description: `Allowed message authentication code (MAC) algorithms

`,
		},

		"port": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Port for SSH service

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port for SSH service

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,

			// Default:          stringdefault.StaticString(`22`),
			Computed: true,
		},

		"client_keepalive_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Enable transmission of keepalives from server to client

    |  Format   |  Description                                     |
    |-----------|--------------------------------------------------|
    |  1-65535  |  Time interval in seconds for keepalive message  |
`,
			Description: `Enable transmission of keepalives from server to client

    |  Format   |  Description                                     |
    |-----------|--------------------------------------------------|
    |  1-65535  |  Time interval in seconds for keepalive message  |
`,
		},

		"vrf": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `VRF instance name

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  default  |  Explicitly start in default VRF  |
    |  txt      |  VRF instance name                |
`,
			Description: `VRF instance name

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  default  |  Explicitly start in default VRF  |
    |  txt      |  VRF instance name                |
`,

			// Default:          stringdefault.StaticString(`default`),
			Computed: true,
		},
	}
}
