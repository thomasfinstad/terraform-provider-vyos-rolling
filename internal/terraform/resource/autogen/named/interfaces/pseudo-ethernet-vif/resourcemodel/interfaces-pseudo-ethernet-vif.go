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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &InterfacesPseudoEthernetVif{}

// InterfacesPseudoEthernetVif describes the resource data model.
type InterfacesPseudoEthernetVif struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafInterfacesPseudoEthernetVifDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesPseudoEthernetVifAddress           types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesPseudoEthernetVifDisableLinkDetect types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesPseudoEthernetVifDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesPseudoEthernetVifEgressQos         types.String `tfsdk:"egress_qos" vyos:"egress-qos,omitempty"`
	LeafInterfacesPseudoEthernetVifIngressQos        types.String `tfsdk:"ingress_qos" vyos:"ingress-qos,omitempty"`
	LeafInterfacesPseudoEthernetVifMac               types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesPseudoEthernetVifMtu               types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesPseudoEthernetVifRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesPseudoEthernetVifVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesPseudoEthernetVifDhcpOptions     *InterfacesPseudoEthernetVifDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesPseudoEthernetVifDhcpvsixOptions *InterfacesPseudoEthernetVifDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesPseudoEthernetVifIP              *InterfacesPseudoEthernetVifIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesPseudoEthernetVifIPvsix          *InterfacesPseudoEthernetVifIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesPseudoEthernetVifMirror          *InterfacesPseudoEthernetVifMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesPseudoEthernetVif) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesPseudoEthernetVif) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesPseudoEthernetVif) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPseudoEthernetVif) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"vif",
		o.SelfIdentifier.Attributes()["vif"].(types.Number).ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesPseudoEthernetVif) GetVyosParentPath() []string {
	return []string{
		"interfaces",

		"pseudo-ethernet",

		o.SelfIdentifier.Attributes()["pseudo_ethernet"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesPseudoEthernetVif) GetVyosNamedParentPath() []string {
	return []string{
		"interfaces",

		"pseudo-ethernet",

		o.SelfIdentifier.Attributes()["pseudo_ethernet"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetVif) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"vif": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Virtual Local Area Network (VLAN) ID

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  0-4094  |  Virtual Local Area Network (VLAN) ID  |
`,
					Description: `Virtual Local Area Network (VLAN) ID

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  0-4094  |  Virtual Local Area Network (VLAN) ID  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				"pseudo_ethernet": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Pseudo Ethernet Interface (Macvlan)

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  pethN   |  Pseudo Ethernet interface name  |
`,
					Description: `Pseudo Ethernet Interface (Macvlan)

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  pethN   |  Pseudo Ethernet interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in pseudo_ethernet, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  pseudo_ethernet, value must match: ^[a-zA-Z0-9-_]*$",
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

		"egress_qos": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VLAN egress QoS

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  txt     |  Format for qos mapping, e.g.: '0:1 1:6 7:6'  |
`,
			Description: `VLAN egress QoS

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  txt     |  Format for qos mapping, e.g.: '0:1 1:6 7:6'  |
`,
		},

		"ingress_qos": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VLAN ingress QoS

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  txt     |  Format for qos mapping, e.g.: '0:1 1:6 7:6'  |
`,
			Description: `VLAN ingress QoS

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  txt     |  Format for qos mapping, e.g.: '0:1 1:6 7:6'  |
`,
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
			Attributes: InterfacesPseudoEthernetVifDhcpOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
			Description: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetVifDhcpvsixOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
			Description: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetVifIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetVifIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetVifMirror{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
			Description: `Mirror ingress/egress packets

`,
		},
	}
}
