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

var _ helpers.VyosTopResourceDataModel = &InterfacesBrIDge{}

// InterfacesBrIDge describes the resource data model.
type InterfacesBrIDge struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafInterfacesBrIDgeAddress           types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesBrIDgeAging             types.Number `tfsdk:"aging" vyos:"aging,omitempty"`
	LeafInterfacesBrIDgeDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesBrIDgeDisableLinkDetect types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesBrIDgeDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesBrIDgeVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafInterfacesBrIDgeMtu               types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesBrIDgeForwardingDelay   types.Number `tfsdk:"forwarding_delay" vyos:"forwarding-delay,omitempty"`
	LeafInterfacesBrIDgeHelloTime         types.Number `tfsdk:"hello_time" vyos:"hello-time,omitempty"`
	LeafInterfacesBrIDgeMac               types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesBrIDgeEnableVlan        types.Bool   `tfsdk:"enable_vlan" vyos:"enable-vlan,omitempty"`
	LeafInterfacesBrIDgeProtocol          types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafInterfacesBrIDgeMaxAge            types.Number `tfsdk:"max_age" vyos:"max-age,omitempty"`
	LeafInterfacesBrIDgePriority          types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafInterfacesBrIDgeStp               types.Bool   `tfsdk:"stp" vyos:"stp,omitempty"`
	LeafInterfacesBrIDgeRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagInterfacesBrIDgeVif bool `tfsdk:"-" vyos:"vif,child"`

	// Nodes
	NodeInterfacesBrIDgeDhcpOptions     *InterfacesBrIDgeDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesBrIDgeDhcpvsixOptions *InterfacesBrIDgeDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesBrIDgeIgmp            *InterfacesBrIDgeIgmp            `tfsdk:"igmp" vyos:"igmp,omitempty"`
	NodeInterfacesBrIDgeIP              *InterfacesBrIDgeIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesBrIDgeIPvsix          *InterfacesBrIDgeIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesBrIDgeMirror          *InterfacesBrIDgeMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
	NodeInterfacesBrIDgeMember          *InterfacesBrIDgeMember          `tfsdk:"member" vyos:"member,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesBrIDge) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesBrIDge) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesBrIDge) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBrIDge) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"bridge",
		o.SelfIdentifier.Attributes()["bridge"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesBrIDge) GetVyosParentPath() []string {
	return []string{
		"interfaces",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesBrIDge) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDge) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"bridge": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `Bridge Interface

    |  Format  |  Description            |
    |----------|-------------------------|
    |  brN     |  Bridge interface name  |
`,
						Description: `Bridge Interface

    |  Format  |  Description            |
    |----------|-------------------------|
    |  brN     |  Bridge interface name  |
`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in bridge, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  bridge, value must match: ^[a-zA-Z0-9-_]*$",
								),
							),
						},
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

		"aging": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `MAC address aging interval

    |  Format      |  Description                                  |
    |--------------|-----------------------------------------------|
    |  0           |  Disable MAC address learning (always flood)  |
    |  10-1000000  |  MAC address aging time in seconds            |
`,
			Description: `MAC address aging interval

    |  Format      |  Description                                  |
    |--------------|-----------------------------------------------|
    |  0           |  Disable MAC address learning (always flood)  |
    |  10-1000000  |  MAC address aging time in seconds            |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
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

		"disable_link_detect": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore link state changes

`,
			Description: `Ignore link state changes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
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

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"forwarding_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Forwarding delay

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  0-200   |  Spanning Tree Protocol forwarding delay in seconds  |
`,
			Description: `Forwarding delay

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  0-200   |  Spanning Tree Protocol forwarding delay in seconds  |
`,

			// Default:          stringdefault.StaticString(`14`),
			Computed: true,
		},

		"hello_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello packet advertisement interval

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  1-10    |  Spanning Tree Protocol hello advertisement interval in seconds  |
`,
			Description: `Hello packet advertisement interval

    |  Format  |  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  1-10    |  Spanning Tree Protocol hello advertisement interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},

		"mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
			Description: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
		},

		"enable_vlan": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable VLAN aware bridge

`,
			Description: `Enable VLAN aware bridge

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol used for service VLAN (default: 802.1ad)

    |  Format   |  Description                                                |
    |-----------|-------------------------------------------------------------|
    |  802.1ad  |  Provider Bridging (IEEE 802.1ad, Q-inQ), ethertype 0x88a8  |
    |  802.1q   |  VLAN-tagged frame (IEEE 802.1q), ethertype 0x8100          |
`,
			Description: `Protocol used for service VLAN (default: 802.1ad)

    |  Format   |  Description                                                |
    |-----------|-------------------------------------------------------------|
    |  802.1ad  |  Provider Bridging (IEEE 802.1ad, Q-inQ), ethertype 0x88a8  |
    |  802.1q   |  VLAN-tagged frame (IEEE 802.1q), ethertype 0x8100          |
`,

			// Default:          stringdefault.StaticString(`802.1q`),
			Computed: true,
		},

		"max_age": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval at which neighbor bridges are removed

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  1-40    |  Bridge maximum aging time in seconds  |
`,
			Description: `Interval at which neighbor bridges are removed

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  1-40    |  Bridge maximum aging time in seconds  |
`,

			// Default:          stringdefault.StaticString(`20`),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Priority for this bridge

    |  Format   |  Description      |
    |-----------|-------------------|
    |  0-65535  |  Bridge priority  |
`,
			Description: `Priority for this bridge

    |  Format   |  Description      |
    |-----------|-------------------|
    |  0-65535  |  Bridge priority  |
`,

			// Default:          stringdefault.StaticString(`32768`),
			Computed: true,
		},

		"stp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable spanning tree protocol

`,
			Description: `Enable spanning tree protocol

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
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

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeDhcpOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
			Description: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeDhcpvsixOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
			Description: `DHCPv6 client settings/options

`,
		},

		"igmp": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeIgmp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Internet Group Management Protocol (IGMP) and Multicast Listener Discovery (MLD) settings

`,
			Description: `Internet Group Management Protocol (IGMP) and Multicast Listener Discovery (MLD) settings

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeMirror{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
			Description: `Mirror ingress/egress packets

`,
		},

		"member": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeMember{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Bridge member interfaces

`,
			Description: `Bridge member interfaces

`,
		},
	}
}
