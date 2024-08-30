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

var _ helpers.VyosTopResourceDataModel = &InterfacesMacsec{}

// InterfacesMacsec describes the resource data model.
type InterfacesMacsec struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafInterfacesMacsecAddress         types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesMacsecDescrIPtion     types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesMacsecDisable         types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesMacsecMtu             types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesMacsecSourceInterface types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`
	LeafInterfacesMacsecRedirect        types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesMacsecVrf             types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesMacsecDhcpOptions     *InterfacesMacsecDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesMacsecDhcpvsixOptions *InterfacesMacsecDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesMacsecIP              *InterfacesMacsecIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesMacsecIPvsix          *InterfacesMacsecIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesMacsecMirror          *InterfacesMacsecMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
	NodeInterfacesMacsecSecURIty        *InterfacesMacsecSecURIty        `tfsdk:"security" vyos:"security,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesMacsec) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesMacsec) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesMacsec) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesMacsec) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"macsec",
		o.SelfIdentifier.Attributes()["macsec"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesMacsec) GetVyosParentPath() []string {
	return []string{
		"interfaces",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesMacsec) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"macsec": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `MACsec Interface (802.1ae)

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  macsecN  |  MACsec interface name  |
`,
					Description: `MACsec Interface (802.1ae)

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  macsecN  |  MACsec interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in macsec, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  macsec, value must match: ^[a-zA-Z0-9-_]*$",
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

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address

    |  Format   |  Description                                   |
    |-----------|------------------------------------------------|
    |  ipv4net  |  IPv4 address and prefix length                |
    |  ipv6net  |  IPv6 address and prefix length                |
    |  dhcp     |  Dynamic Host Configuration Protocol           |
    |  dhcpv6   |  Dynamic Host Configuration Protocol for IPv6  |
`,
			Description: `IP address

    |  Format   |  Description                                   |
    |-----------|------------------------------------------------|
    |  ipv4net  |  IPv4 address and prefix length                |
    |  ipv6net  |  IPv6 address and prefix length                |
    |  dhcp     |  Dynamic Host Configuration Protocol           |
    |  dhcpv6   |  Dynamic Host Configuration Protocol for IPv6  |
`,
		},

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

			// Default:          stringdefault.StaticString(`1460`),
			Computed: true,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Physical interface the traffic will go through

    |  Format     |  Description                                     |
    |-------------|--------------------------------------------------|
    |  interface  |  Physical interface used for traffic forwarding  |
`,
			Description: `Physical interface the traffic will go through

    |  Format     |  Description                                     |
    |-------------|--------------------------------------------------|
    |  interface  |  Physical interface used for traffic forwarding  |
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

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecDhcpOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
			Description: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecDhcpvsixOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
			Description: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecMirror{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
			Description: `Mirror ingress/egress packets

`,
		},

		"security": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecSecURIty{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Security/Encryption Settings

`,
			Description: `Security/Encryption Settings

`,
		},
	}
}
