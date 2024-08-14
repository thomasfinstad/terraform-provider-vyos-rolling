// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsBgpAddressFamilyLtwovpnEvpnVni{}

// ProtocolsBgpAddressFamilyLtwovpnEvpnVni describes the resource data model.
type ProtocolsBgpAddressFamilyLtwovpnEvpnVni struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseDefaultGw types.Bool   `tfsdk:"advertise_default_gw" vyos:"advertise-default-gw,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseSviIP     types.Bool   `tfsdk:"advertise_svi_ip" vyos:"advertise-svi-ip,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnVniRd                 types.String `tfsdk:"rd" vyos:"rd,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget *ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget `tfsdk:"route_target" vyos:"route-target,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"vni",
		o.SelfIdentifier.Attributes()["vni"].(types.Number).ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"bgp",

		"address-family",

		"l2vpn-evpn",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyLtwovpnEvpnVni) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"vni": schema.NumberAttribute{
						Required: true,
						MarkdownDescription: `VXLAN Network Identifier

    |  Format      |  Description  |
    |--------------|---------------|
    |  1-16777215  |  VNI number   |
`,
						Description: `VXLAN Network Identifier

    |  Format      |  Description  |
    |--------------|---------------|
    |  1-16777215  |  VNI number   |
`,
						PlanModifiers: []planmodifier.Number{
							numberplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

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

		// Nodes

		"route_target": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route Target

`,
			Description: `Route Target

`,
		},
	}
}