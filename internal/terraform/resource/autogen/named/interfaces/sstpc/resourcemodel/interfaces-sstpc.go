// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &InterfacesSstpc{}

// InterfacesSstpc describes the resource data model.
type InterfacesSstpc struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafInterfacesSstpcDescrIPtion          types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesSstpcDisable              types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesSstpcNoDefaultRoute       types.Bool   `tfsdk:"no_default_route" vyos:"no-default-route,omitempty"`
	LeafInterfacesSstpcDefaultRouteDistance types.Number `tfsdk:"default_route_distance" vyos:"default-route-distance,omitempty"`
	LeafInterfacesSstpcNoPeerDNS            types.Bool   `tfsdk:"no_peer_dns" vyos:"no-peer-dns,omitempty"`
	LeafInterfacesSstpcMtu                  types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesSstpcServer               types.String `tfsdk:"server" vyos:"server,omitempty"`
	LeafInterfacesSstpcPort                 types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafInterfacesSstpcVrf                  types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesSstpcAuthentication *InterfacesSstpcAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`
	NodeInterfacesSstpcSsl            *InterfacesSstpcSsl            `tfsdk:"ssl" vyos:"ssl,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesSstpc) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesSstpc) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesSstpc) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesSstpc) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"sstpc",
		o.SelfIdentifier.Attributes()["sstpc"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesSstpc) GetVyosParentPath() []string {
	return []string{
		"interfaces",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesSstpc) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesSstpc) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"sstpc": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Secure Socket Tunneling Protocol (SSTP) client Interface

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  sstpcN  |  Secure Socket Tunneling Protocol interface name  |
`,
					Description: `Secure Socket Tunneling Protocol (SSTP) client Interface

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  sstpcN  |  Secure Socket Tunneling Protocol interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in sstpc, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  sstpc, value must match: ^[a-zA-Z0-9-_]*$",
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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Description: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_default_route": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not install default route to system

`,
			Description: `Do not install default route to system

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"default_route_distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for installed default route

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  1-255   |  Distance for the default route from DHCP server  |
`,
			Description: `Distance for installed default route

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  1-255   |  Distance for the default route from DHCP server  |
`,

			// Default:          stringdefault.StaticString(`210`),
			Computed: true,
		},

		"no_peer_dns": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not use DNS servers provided by the peer

`,
			Description: `Do not use DNS servers provided by the peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  68-1500  |  Maximum Transmission Unit in byte  |
`,
			Description: `Maximum Transmission Unit (MTU)

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  68-1500  |  Maximum Transmission Unit in byte  |
`,

			// Default:          stringdefault.StaticString(`1452`),
			Computed: true,
		},

		"server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remote server to connect to

    |  Format    |  Description           |
    |------------|------------------------|
    |  ipv4      |  Server IPv4 address   |
    |  hostname  |  Server hostname/FQDN  |
`,
			Description: `Remote server to connect to

    |  Format    |  Description           |
    |------------|------------------------|
    |  ipv4      |  Server IPv4 address   |
    |  hostname  |  Server hostname/FQDN  |
`,
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

			// Default:          stringdefault.StaticString(`443`),
			Computed: true,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
		},

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: InterfacesSstpcAuthentication{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
			Description: `Authentication settings

`,
		},

		"ssl": schema.SingleNestedAttribute{
			Attributes: InterfacesSstpcSsl{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Secure Sockets Layer (SSL) configuration

`,
			Description: `Secure Sockets Layer (SSL) configuration

`,
		},
	}
}
