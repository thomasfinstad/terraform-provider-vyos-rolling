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
var _ helpers.VyosTopResourceDataModel = &ProtocolsBgpAddressFamilyLtwovpnEvpn{}

// ProtocolsBgpAddressFamilyLtwovpnEvpn describes the resource data model.
type ProtocolsBgpAddressFamilyLtwovpnEvpn struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseAllVni    types.Bool   `tfsdk:"advertise_all_vni" vyos:"advertise-all-vni,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseDefaultGw types.Bool   `tfsdk:"advertise_default_gw" vyos:"advertise-default-gw,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseSviIP     types.Bool   `tfsdk:"advertise_svi_ip" vyos:"advertise-svi-ip,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnRd                 types.String `tfsdk:"rd" vyos:"rd,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnAdvertisePIP       types.String `tfsdk:"advertise_pip" vyos:"advertise-pip,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnRtAutoDerive       types.Bool   `tfsdk:"rt_auto_derive" vyos:"rt-auto-derive,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnDisableEadEviRx    types.Bool   `tfsdk:"disable_ead_evi_rx" vyos:"disable-ead-evi-rx,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnDisableEadEviTx    types.Bool   `tfsdk:"disable_ead_evi_tx" vyos:"disable-ead-evi-tx,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsBgpAddressFamilyLtwovpnEvpnVni bool `tfsdk:"-" vyos:"vni,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise        bool `tfsdk:"-" vyos:"advertise,child"`
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget      bool `tfsdk:"-" vyos:"route-target,child"`
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnDefaultOriginate bool `tfsdk:"-" vyos:"default-originate,child"`
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag        bool `tfsdk:"-" vyos:"ead-es-frag,child"`
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget bool `tfsdk:"-" vyos:"ead-es-route-target,child"`
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnFlooding         bool `tfsdk:"-" vyos:"flooding,child"`
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnMacVrf           bool `tfsdk:"-" vyos:"mac-vrf,child"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpn) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpn) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpn) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpn) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"l2vpn-evpn",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpn) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"bgp",

		"address-family",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpn) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyLtwovpnEvpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"advertise_all_vni": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise All local VNIs

`,
			Description: `Advertise All local VNIs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"advertise_default_gw": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise All default g/w mac-ip routes in EVPN

`,
			Description: `Advertise All default g/w mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"advertise_svi_ip": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise svi mac-ip routes in EVPN

`,
			Description: `Advertise svi mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
			Description: `Route Distinguisher

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
		},

		"advertise_pip": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `EVPN system primary IP

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4    |  IP address   |
`,
			Description: `EVPN system primary IP

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4    |  IP address   |
`,
		},

		"rt_auto_derive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Auto derivation of Route Target (RFC8365)

`,
			Description: `Auto derivation of Route Target (RFC8365)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_ead_evi_rx": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Activate PE on EAD-ES even if EAD-EVI is not received

`,
			Description: `Activate PE on EAD-ES even if EAD-EVI is not received

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_ead_evi_tx": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not advertise EAD-EVI for local ESs

`,
			Description: `Do not advertise EAD-EVI for local ESs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}