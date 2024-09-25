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
var _ helpers.VyosTopResourceDataModel = &VpnLtwotpRemoteAccessAuthentication{}

// VpnLtwotpRemoteAccessAuthentication describes the resource data model.
type VpnLtwotpRemoteAccessAuthentication struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnLtwotpRemoteAccessAuthenticationMode      types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafVpnLtwotpRemoteAccessAuthenticationProtocols types.List   `tfsdk:"protocols" vyos:"protocols,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeVpnLtwotpRemoteAccessAuthenticationLocalUsers bool `tfsdk:"-" vyos:"local-users,child"`
	ExistsNodeVpnLtwotpRemoteAccessAuthenticationRadius     bool `tfsdk:"-" vyos:"radius,child"`
}

// SetID configures the resource ID
func (o *VpnLtwotpRemoteAccessAuthentication) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnLtwotpRemoteAccessAuthentication) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnLtwotpRemoteAccessAuthentication) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnLtwotpRemoteAccessAuthentication) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"authentication",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnLtwotpRemoteAccessAuthentication) GetVyosParentPath() []string {
	return []string{
		"vpn",

		"l2tp",

		"remote-access",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *VpnLtwotpRemoteAccessAuthentication) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnLtwotpRemoteAccessAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication mode used by this server

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  local   |  Use local username/password configuration  |
    |  radius  |  Use RADIUS server for user autentication   |
    |  noauth  |  Authentication disabled                    |
`,
			Description: `Authentication mode used by this server

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  local   |  Use local username/password configuration  |
    |  radius  |  Use RADIUS server for user autentication   |
    |  noauth  |  Authentication disabled                    |
`,

			// Default:          stringdefault.StaticString(`local`),
			Computed: true,
		},

		"protocols": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Authentication protocol for remote access peer

    |  Format     |  Description                                                                                      |
    |-------------|---------------------------------------------------------------------------------------------------|
    |  pap        |  Authentication via PAP (Password Authentication Protocol)                                        |
    |  chap       |  Authentication via CHAP (Challenge Handshake Authentication Protocol)                            |
    |  mschap     |  Authentication via MS-CHAP (Microsoft Challenge Handshake Authentication Protocol)               |
    |  mschap-v2  |  Authentication via MS-CHAPv2 (Microsoft Challenge Handshake Authentication Protocol, version 2)  |
`,
			Description: `Authentication protocol for remote access peer

    |  Format     |  Description                                                                                      |
    |-------------|---------------------------------------------------------------------------------------------------|
    |  pap        |  Authentication via PAP (Password Authentication Protocol)                                        |
    |  chap       |  Authentication via CHAP (Challenge Handshake Authentication Protocol)                            |
    |  mschap     |  Authentication via MS-CHAP (Microsoft Challenge Handshake Authentication Protocol)               |
    |  mschap-v2  |  Authentication via MS-CHAPv2 (Microsoft Challenge Handshake Authentication Protocol, version 2)  |
`,

			// Default:          stringdefault.StaticString(`pap chap mschap mschap-v2`),
			Computed: true,
		},
	}
}
