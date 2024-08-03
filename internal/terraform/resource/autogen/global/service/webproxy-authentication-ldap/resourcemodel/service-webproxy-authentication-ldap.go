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
var _ helpers.VyosTopResourceDataModel = &ServiceWebproxyAuthenticationLdap{}

// ServiceWebproxyAuthenticationLdap describes the resource data model.
type ServiceWebproxyAuthenticationLdap struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceWebproxyAuthenticationLdapBaseDn               types.String `tfsdk:"base_dn" vyos:"base-dn,omitempty"`
	LeafServiceWebproxyAuthenticationLdapBindDn               types.String `tfsdk:"bind_dn" vyos:"bind-dn,omitempty"`
	LeafServiceWebproxyAuthenticationLdapFilterExpression     types.String `tfsdk:"filter_expression" vyos:"filter-expression,omitempty"`
	LeafServiceWebproxyAuthenticationLdapPassword             types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafServiceWebproxyAuthenticationLdapPersistentConnection types.Bool   `tfsdk:"persistent_connection" vyos:"persistent-connection,omitempty"`
	LeafServiceWebproxyAuthenticationLdapPort                 types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafServiceWebproxyAuthenticationLdapServer               types.String `tfsdk:"server" vyos:"server,omitempty"`
	LeafServiceWebproxyAuthenticationLdapUseSsl               types.Bool   `tfsdk:"use_ssl" vyos:"use-ssl,omitempty"`
	LeafServiceWebproxyAuthenticationLdapUsernameAttribute    types.String `tfsdk:"username_attribute" vyos:"username-attribute,omitempty"`
	LeafServiceWebproxyAuthenticationLdapVersion              types.String `tfsdk:"version" vyos:"version,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceWebproxyAuthenticationLdap) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceWebproxyAuthenticationLdap) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceWebproxyAuthenticationLdap) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceWebproxyAuthenticationLdap) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"ldap",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceWebproxyAuthenticationLdap) GetVyosParentPath() []string {
	return []string{
		"service",

		"webproxy",

		"authentication",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ServiceWebproxyAuthenticationLdap) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceWebproxyAuthenticationLdap) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"base_dn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `LDAP Base DN to search

`,
			Description: `LDAP Base DN to search

`,
		},

		"bind_dn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `LDAP DN used to bind to server

`,
			Description: `LDAP DN used to bind to server

`,
		},

		"filter_expression": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Filter expression to perform LDAP search with

`,
			Description: `Filter expression to perform LDAP search with

`,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `LDAP password to bind with

`,
			Description: `LDAP password to bind with

`,
		},

		"persistent_connection": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use persistent LDAP connection

`,
			Description: `Use persistent LDAP connection

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,

			// Default:          stringdefault.StaticString(`389`),
			Computed: true,
		},

		"server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `LDAP server to use

`,
			Description: `LDAP server to use

`,
		},

		"use_ssl": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use SSL/TLS for LDAP connection

`,
			Description: `Use SSL/TLS for LDAP connection

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"username_attribute": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `LDAP username attribute

`,
			Description: `LDAP username attribute

`,
		},

		"version": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `LDAP protocol version

    |  Format  |  Description              |
    |----------|---------------------------|
    |  2       |  LDAP protocol version 2  |
    |  3       |  LDAP protocol version 2  |
`,
			Description: `LDAP protocol version

    |  Format  |  Description              |
    |----------|---------------------------|
    |  2       |  LDAP protocol version 2  |
    |  3       |  LDAP protocol version 2  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},
	}
}
