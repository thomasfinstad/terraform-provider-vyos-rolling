/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (vni) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (vni) */
// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniSelfIdentifier used by TagNodes to keep the id info
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (vni) */

	VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni types.Number `tfsdk:"vni" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (l2vpn-evpn) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (address-family) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (bgp) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (protocols) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (name) */

	VrfName types.String `tfsdk:"name" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (vrf) */
}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseDefaultGw types.Bool   `tfsdk:"advertise_default_gw" vyos:"advertise-default-gw,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniAdvertiseSviIP     types.Bool   `tfsdk:"advertise_svi_ip" vyos:"advertise-svi-ip,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRd                 types.String `tfsdk:"rd" vyos:"rd,omitempty"`

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget `tfsdk:"route_target" vyos:"route-target,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"vni",
		o.SelfIdentifier.VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (l2vpn-evpn) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (address-family) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (bgp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (name) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vrf) */
		"vrf", // Node

		"name",
		o.SelfIdentifier.VrfName.ValueString(),

		"protocols", // Node

		"bgp", // Node

		"address-family", // Node

		"l2vpn-evpn", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (l2vpn-evpn) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (address-family) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (bgp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (protocols) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (name) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (name) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vrf) */
		"vrf", // Node

		"name",
		o.SelfIdentifier.VrfName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
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

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (l2vpn-evpn) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (address-family) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (bgp) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (protocols) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (name) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (vrf) */

				"name": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
					Description: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  name, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"advertise_default_gw":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (advertise-default-gw) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise All default g/w mac-ip routes in EVPN

`,
			Description: `Advertise All default g/w mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"advertise_svi_ip":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (advertise-svi-ip) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise svi mac-ip routes in EVPN

`,
			Description: `Advertise svi mac-ip routes in EVPN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rd":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (rd) */
		schema.StringAttribute{
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

		// TagNodes

		// Nodes

		"route_target": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVniRouteTarget{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route Target

`,
			Description: `Route Target

`,
		},
	}
}
