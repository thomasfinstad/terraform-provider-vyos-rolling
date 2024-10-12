// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceDNSDynamicName{}

// ServiceDNSDynamicName describes the resource data model.
type ServiceDNSDynamicName struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDNSDynamicNameDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafServiceDNSDynamicNameProtocol    types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafServiceDNSDynamicNameIPVersion   types.String `tfsdk:"ip_version" vyos:"ip-version,omitempty"`
	LeafServiceDNSDynamicNameHostName    types.List   `tfsdk:"host_name" vyos:"host-name,omitempty"`
	LeafServiceDNSDynamicNameServer      types.String `tfsdk:"server" vyos:"server,omitempty"`
	LeafServiceDNSDynamicNameZone        types.String `tfsdk:"zone" vyos:"zone,omitempty"`
	LeafServiceDNSDynamicNameUsername    types.String `tfsdk:"username" vyos:"username,omitempty"`
	LeafServiceDNSDynamicNamePassword    types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafServiceDNSDynamicNameKey         types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafServiceDNSDynamicNameTTL         types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`
	LeafServiceDNSDynamicNameWaitTime    types.Number `tfsdk:"wait_time" vyos:"wait-time,omitempty"`
	LeafServiceDNSDynamicNameExpiryTime  types.Number `tfsdk:"expiry_time" vyos:"expiry-time,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeServiceDNSDynamicNameAddress *ServiceDNSDynamicNameAddress `tfsdk:"address" vyos:"address,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceDNSDynamicName) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDNSDynamicName) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDNSDynamicName) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSDynamicName) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"name",
		o.SelfIdentifier.Attributes()["name"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDNSDynamicName) GetVyosParentPath() []string {
	return []string{
		"service",

		"dns",

		"dynamic",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDNSDynamicName) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSDynamicName) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Dynamic DNS configuration

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Dynamic DNS service name  |
`,
					Description: `Dynamic DNS configuration

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Dynamic DNS service name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in name, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  name, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description": schema.StringAttribute{
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

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ddclient protocol used for Dynamic DNS service

`,
			Description: `ddclient protocol used for Dynamic DNS service

`,
		},

		"ip_version": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address version to use

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  _ipv4   |  Use only IPv4 address           |
    |  _ipv6   |  Use only IPv6 address           |
    |  both    |  Use both IPv4 and IPv6 address  |
`,
			Description: `IP address version to use

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  _ipv4   |  Use only IPv4 address           |
    |  _ipv6   |  Use only IPv6 address           |
    |  both    |  Use both IPv4 and IPv6 address  |
`,

			// Default:          stringdefault.StaticString(`ipv4`),
			Computed: true,
		},

		"host_name": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Hostname to register with Dynamic DNS service

`,
			Description: `Hostname to register with Dynamic DNS service

`,
		},

		"server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remote Dynamic DNS server to send updates to

    |  Format    |  Description                                       |
    |------------|----------------------------------------------------|
    |  ipv4      |  IPv4 address of the remote server                 |
    |  ipv6      |  IPv6 address of the remote server                 |
    |  hostname  |  Fully qualified domain name of the remote server  |
`,
			Description: `Remote Dynamic DNS server to send updates to

    |  Format    |  Description                                       |
    |------------|----------------------------------------------------|
    |  ipv4      |  IPv4 address of the remote server                 |
    |  ipv6      |  IPv6 address of the remote server                 |
    |  hostname  |  Fully qualified domain name of the remote server  |
`,
		},

		"zone": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DNS zone to be updated

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  Name of DNS zone  |
`,
			Description: `DNS zone to be updated

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  Name of DNS zone  |
`,
		},

		"username": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Username     |
`,
			Description: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Username     |
`,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Password     |
`,
			Description: `Password used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Password     |
`,
		},

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `File containing TSIG authentication key for RFC2136 nsupdate on remote DNS server

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  filename  |  File in /config/auth directory  |
`,
			Description: `File containing TSIG authentication key for RFC2136 nsupdate on remote DNS server

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  filename  |  File in /config/auth directory  |
`,
		},

		"ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time-to-live (TTL)

    |  Format        |  Description     |
    |----------------|------------------|
    |  0-2147483647  |  TTL in seconds  |
`,
			Description: `Time-to-live (TTL)

    |  Format        |  Description     |
    |----------------|------------------|
    |  0-2147483647  |  TTL in seconds  |
`,
		},

		"wait_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time in seconds to wait between update attempts

    |  Format    |  Description      |
    |------------|-------------------|
    |  60-86400  |  Time in seconds  |
`,
			Description: `Time in seconds to wait between update attempts

    |  Format    |  Description      |
    |------------|-------------------|
    |  60-86400  |  Time in seconds  |
`,
		},

		"expiry_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time in seconds for the hostname to be marked expired in cache

    |  Format       |  Description      |
    |---------------|-------------------|
    |  300-2160000  |  Time in seconds  |
`,
			Description: `Time in seconds for the hostname to be marked expired in cache

    |  Format       |  Description      |
    |---------------|-------------------|
    |  300-2160000  |  Time in seconds  |
`,
		},

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: ServiceDNSDynamicNameAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Obtain IP address to send Dynamic DNS update for

`,
			Description: `Obtain IP address to send Dynamic DNS update for

`,
		},
	}
}
