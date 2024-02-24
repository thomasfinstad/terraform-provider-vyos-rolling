// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpAddressFamilyLtwovpnEvpn describes the resource data model.
type ProtocolsBgpAddressFamilyLtwovpnEvpn struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseAllVni    types.Bool   `tfsdk:"advertise_all_vni" vyos:"advertise-all-vni,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseDefaultGw types.Bool   `tfsdk:"advertise_default_gw" vyos:"advertise-default-gw,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseSviIP     types.Bool   `tfsdk:"advertise_svi_ip" vyos:"advertise-svi-ip,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnRd                 types.String `tfsdk:"rd" vyos:"rd,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnAdvertisePIP       types.String `tfsdk:"advertise_pip" vyos:"advertise-pip,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnRtAutoDerive       types.Bool   `tfsdk:"rt_auto_derive" vyos:"rt-auto-derive,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsBgpAddressFamilyLtwovpnEvpnVni bool `tfsdk:"-" vyos:"vni,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnAdvertise   bool `tfsdk:"-" vyos:"advertise,ignore,omitempty"`
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget bool `tfsdk:"-" vyos:"route-target,ignore,omitempty"`
	ExistsNodeProtocolsBgpAddressFamilyLtwovpnEvpnFlooding    bool `tfsdk:"-" vyos:"flooding,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpn) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpn) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"bgp",

		"address-family",

		"l2vpn-evpn",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyLtwovpnEvpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"advertise_all_vni": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise All local VNIs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"advertise_default_gw": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise All default g/w mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"advertise_svi_ip": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise svi mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  &emsp; |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |

`,
		},

		"advertise_pip": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `EVPN system primary IP

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IP address  |

`,
		},

		"rt_auto_derive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Auto derivation of Route Target (RFC8365)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}