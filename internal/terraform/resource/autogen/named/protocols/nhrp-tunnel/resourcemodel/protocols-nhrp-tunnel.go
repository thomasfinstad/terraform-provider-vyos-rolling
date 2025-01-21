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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (tunnel) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsNhrpTunnel{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (tunnel) */
// ProtocolsNhrpTunnelSelfIdentifier used by TagNodes to keep the id info
type ProtocolsNhrpTunnelSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (tunnel) */

	ProtocolsNhrpTunnel types.String `tfsdk:"tunnel" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (nhrp) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (protocols) */
}

// ProtocolsNhrpTunnel describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ProtocolsNhrpTunnel struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *ProtocolsNhrpTunnelSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafProtocolsNhrpTunnelMulticast            types.List   `tfsdk:"multicast" vyos:"multicast,omitempty"`
	LeafProtocolsNhrpTunnelRegistrationNoUnique types.Bool   `tfsdk:"registration_no_unique" vyos:"registration-no-unique,omitempty"`
	LeafProtocolsNhrpTunnelAuthentication       types.String `tfsdk:"authentication" vyos:"authentication,omitempty"`
	LeafProtocolsNhrpTunnelHoldtime             types.Number `tfsdk:"holdtime" vyos:"holdtime,omitempty"`
	LeafProtocolsNhrpTunnelRedirect             types.Bool   `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafProtocolsNhrpTunnelShortcut             types.Bool   `tfsdk:"shortcut" vyos:"shortcut,omitempty"`
	LeafProtocolsNhrpTunnelMtu                  types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafProtocolsNhrpTunnelNetworkID            types.String `tfsdk:"network_id" vyos:"network-id,omitempty"`

	// TagNodes

	// Nodes

	// Ignoring Node `ProtocolsNhrpTunnelMap`.

	// Ignoring Node `ProtocolsNhrpTunnelNhs`.
}

// SetID configures the resource ID
func (o *ProtocolsNhrpTunnel) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsNhrpTunnel) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsNhrpTunnel) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsNhrpTunnel) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"tunnel",
		o.SelfIdentifier.ProtocolsNhrpTunnel.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsNhrpTunnel) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (nhrp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */
		"protocols", // Node

		"nhrp", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsNhrpTunnel) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (nhrp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (protocols) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsNhrpTunnel) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `Tunnel for NHRP

    |  Format  |  Description       |
    |----------|--------------------|
    |  tunN    |  NHRP tunnel name  |
`,
					Description: `Tunnel for NHRP

    |  Format  |  Description       |
    |----------|--------------------|
    |  tunN    |  NHRP tunnel name  |
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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  tunnel, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (nhrp) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (protocols) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"multicast":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (multicast) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Map multicast to NBMA

    |  Format   |  Description                         |
    |-----------|--------------------------------------|
    |  ipv4     |  Set the IP address to map(IP|FQDN)  |
    |  dynamic  |  NBMA address is learnt dynamically  |
`,
			Description: `Map multicast to NBMA

    |  Format   |  Description                         |
    |-----------|--------------------------------------|
    |  ipv4     |  Set the IP address to map(IP|FQDN)  |
    |  dynamic  |  NBMA address is learnt dynamically  |
`,
		},

		"registration_no_unique":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (registration-no-unique) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Don't set unique flag

`,
			Description: `Don't set unique flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"authentication":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (authentication) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NHRP authentication

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  txt     |  Pass phrase for NHRP authentication  |
`,
			Description: `NHRP authentication

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  txt     |  Pass phrase for NHRP authentication  |
`,
		},

		"holdtime":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (holdtime) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Holding time in seconds

    |  Format   |  Description       |
    |-----------|--------------------|
    |  1-65000  |  ring buffer size  |
`,
			Description: `Holding time in seconds

    |  Format   |  Description       |
    |-----------|--------------------|
    |  1-65000  |  ring buffer size  |
`,
		},

		"redirect":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (redirect) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable sending of Cisco style NHRP Traffic Indication packets

`,
			Description: `Enable sending of Cisco style NHRP Traffic Indication packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"shortcut":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (shortcut) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable creation of shortcut routes. A received NHRP Traffic Indication will trigger the resolution and establishment of a shortcut route

`,
			Description: `Enable creation of shortcut routes. A received NHRP Traffic Indication will trigger the resolution and establishment of a shortcut route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mtu":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mtu) */
		schema.NumberAttribute{
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

		"network_id":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (network-id) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NHRP network id

    |  Format          |  Description      |
    |------------------|-------------------|
    |  <1-4294967295>  |  NHRP network id  |
`,
			Description: `NHRP network id

    |  Format          |  Description      |
    |------------------|-------------------|
    |  <1-4294967295>  |  NHRP network id  |
`,
		},

		// TagNodes

		// Nodes

	}
}
