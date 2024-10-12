// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &InterfacesTunnel{}

// InterfacesTunnel describes the resource data model.
type InterfacesTunnel struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafInterfacesTunnelDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesTunnelAddress           types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesTunnelDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesTunnelDisableLinkDetect types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesTunnelMtu               types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesTunnelSourceAddress     types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafInterfacesTunnelRemote            types.String `tfsdk:"remote" vyos:"remote,omitempty"`
	LeafInterfacesTunnelSourceInterface   types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`
	LeafInterfacesTunnelSixrdPrefix       types.String `tfsdk:"_6rd_prefix" vyos:"6rd-prefix,omitempty"`
	LeafInterfacesTunnelSixrdRelayPrefix  types.String `tfsdk:"_6rd_relay_prefix" vyos:"6rd-relay-prefix,omitempty"`
	LeafInterfacesTunnelEncapsulation     types.String `tfsdk:"encapsulation" vyos:"encapsulation,omitempty"`
	LeafInterfacesTunnelEnableMulticast   types.Bool   `tfsdk:"enable_multicast" vyos:"enable-multicast,omitempty"`
	LeafInterfacesTunnelVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafInterfacesTunnelRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesTunnelIP         *InterfacesTunnelIP         `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesTunnelIPvsix     *InterfacesTunnelIPvsix     `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesTunnelMirror     *InterfacesTunnelMirror     `tfsdk:"mirror" vyos:"mirror,omitempty"`
	NodeInterfacesTunnelParameters *InterfacesTunnelParameters `tfsdk:"parameters" vyos:"parameters,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesTunnel) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesTunnel) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesTunnel) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesTunnel) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"tunnel",
		o.SelfIdentifier.Attributes()["tunnel"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesTunnel) GetVyosParentPath() []string {
	return []string{
		"interfaces",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesTunnel) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesTunnel) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"tunnel": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Tunnel interface

    |  Format  |  Description            |
    |----------|-------------------------|
    |  tunN    |  Tunnel interface name  |
`,
					Description: `Tunnel interface

    |  Format  |  Description            |
    |----------|-------------------------|
    |  tunN    |  Tunnel interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in tunnel, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  tunnel, value must match: ^[a-zA-Z0-9-_]*$",
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

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `IP address

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6net  |  IPv6 address and prefix length  |
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

		"disable_link_detect": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore link state changes

`,
			Description: `Ignore link state changes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  68-16000  |  Maximum Transmission Unit in byte  |
`,
			Description: `Maximum Transmission Unit (MTU)

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  68-16000  |  Maximum Transmission Unit in byte  |
`,

			// Default:          stringdefault.StaticString(`1476`),
			Computed: true,
		},

		"source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source IP address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
    |  ipv6    |  IPv6 source address  |
`,
			Description: `Source IP address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
    |  ipv6    |  IPv6 source address  |
`,
		},

		"remote": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Tunnel remote address

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  Tunnel remote IPv4 address  |
    |  ipv6    |  Tunnel remote IPv6 address  |
`,
			Description: `Tunnel remote address

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  Tunnel remote IPv4 address  |
    |  ipv6    |  Tunnel remote IPv6 address  |
`,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface used to establish connection

    |  Format     |  Description     |
    |-------------|------------------|
    |  interface  |  Interface name  |
`,
			Description: `Interface used to establish connection

    |  Format     |  Description     |
    |-------------|------------------|
    |  interface  |  Interface name  |
`,
		},

		"_6rd_prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `6rd network prefix

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv6    |  IPv6 address and prefix length  |
`,
			Description: `6rd network prefix

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv6    |  IPv6 address and prefix length  |
`,
		},

		"_6rd_relay_prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `6rd relay prefix

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  ipv4net  |  IPv4 prefix of interface for 6rd  |
`,
			Description: `6rd relay prefix

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  ipv4net  |  IPv4 prefix of interface for 6rd  |
`,
		},

		"encapsulation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encapsulation of this tunnel interface

    |  Format     |  Description                                           |
    |-------------|--------------------------------------------------------|
    |  erspan     |  Encapsulated Remote Switched Port Analyzer            |
    |  gre        |  Generic Routing Encapsulation (network layer)         |
    |  gretap     |  Generic Routing Encapsulation (datalink layer)        |
    |  ip6erspan  |  Encapsulated Remote Switched Port Analyzer over IPv6  |
    |  ip6gre     |  GRE over IPv6 (network layer)                         |
    |  ip6gretap  |  GRE over IPv6 (datalink layer)                        |
    |  ip6ip6     |  IPv6 in IPv6 encapsulation                            |
    |  ipip       |  IPv4 in IPv4 encapsulation                            |
    |  ipip6      |  IPv4 in IP6 encapsulation                             |
    |  sit        |  Simple Internet Transition (IPv6 in IPv4)             |
`,
			Description: `Encapsulation of this tunnel interface

    |  Format     |  Description                                           |
    |-------------|--------------------------------------------------------|
    |  erspan     |  Encapsulated Remote Switched Port Analyzer            |
    |  gre        |  Generic Routing Encapsulation (network layer)         |
    |  gretap     |  Generic Routing Encapsulation (datalink layer)        |
    |  ip6erspan  |  Encapsulated Remote Switched Port Analyzer over IPv6  |
    |  ip6gre     |  GRE over IPv6 (network layer)                         |
    |  ip6gretap  |  GRE over IPv6 (datalink layer)                        |
    |  ip6ip6     |  IPv6 in IPv6 encapsulation                            |
    |  ipip       |  IPv4 in IPv4 encapsulation                            |
    |  ipip6      |  IPv4 in IP6 encapsulation                             |
    |  sit        |  Simple Internet Transition (IPv6 in IPv4)             |
`,
		},

		"enable_multicast": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable multicast operation over tunnel

`,
			Description: `Enable multicast operation over tunnel

`,
			Default:  booldefault.StaticBool(false),
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

		"redirect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
			Description: `Redirect incoming packet to destination

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
		},

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelMirror{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
			Description: `Mirror ingress/egress packets

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelParameters{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Tunnel parameters

`,
			Description: `Tunnel parameters

`,
		},
	}
}
