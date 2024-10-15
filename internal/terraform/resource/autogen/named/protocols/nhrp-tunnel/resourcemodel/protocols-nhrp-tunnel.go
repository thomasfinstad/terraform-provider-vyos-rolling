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

var _ helpers.VyosTopResourceDataModel = &ProtocolsNhrpTunnel{}

// ProtocolsNhrpTunnel describes the resource data model.
type ProtocolsNhrpTunnel struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsNhrpTunnelCiscoAuthentication types.String `tfsdk:"cisco_authentication" vyos:"cisco-authentication,omitempty"`
	LeafProtocolsNhrpTunnelHoldingTime         types.String `tfsdk:"holding_time" vyos:"holding-time,omitempty"`
	LeafProtocolsNhrpTunnelMulticast           types.String `tfsdk:"multicast" vyos:"multicast,omitempty"`
	LeafProtocolsNhrpTunnelNonCaching          types.Bool   `tfsdk:"non_caching" vyos:"non-caching,omitempty"`
	LeafProtocolsNhrpTunnelRedirect            types.Bool   `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafProtocolsNhrpTunnelShortcutDestination types.Bool   `tfsdk:"shortcut_destination" vyos:"shortcut-destination,omitempty"`
	LeafProtocolsNhrpTunnelShortcut            types.Bool   `tfsdk:"shortcut" vyos:"shortcut,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagProtocolsNhrpTunnelDynamicMap bool `tfsdk:"-" vyos:"dynamic-map,child"`

	ExistsTagProtocolsNhrpTunnelMap bool `tfsdk:"-" vyos:"map,child"`

	ExistsTagProtocolsNhrpTunnelShortcutTarget bool `tfsdk:"-" vyos:"shortcut-target,child"`

	// Nodes
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
		o.SelfIdentifier.Attributes()["tunnel"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsNhrpTunnel) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"nhrp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsNhrpTunnel) GetVyosNamedParentPath() []string {
	return []string{}
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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  tunnel, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"cisco_authentication": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pass phrase for cisco authentication

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  txt     |  Pass phrase for cisco authentication  |
`,
			Description: `Pass phrase for cisco authentication

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  txt     |  Pass phrase for cisco authentication  |
`,
		},

		"holding_time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Holding time in seconds

`,
			Description: `Holding time in seconds

`,
		},

		"multicast": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set multicast for NHRP

`,
			Description: `Set multicast for NHRP

`,
		},

		"non_caching": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `This can be used to reduce memory consumption on big NBMA subnets

`,
			Description: `This can be used to reduce memory consumption on big NBMA subnets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"redirect": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable sending of Cisco style NHRP Traffic Indication packets

`,
			Description: `Enable sending of Cisco style NHRP Traffic Indication packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"shortcut_destination": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `This instructs opennhrp to reply with authorative answers on NHRP Resolution Requests destined to addresses in this interface

`,
			Description: `This instructs opennhrp to reply with authorative answers on NHRP Resolution Requests destined to addresses in this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"shortcut": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable creation of shortcut routes. A received NHRP Traffic Indication will trigger the resolution and establishment of a shortcut route

`,
			Description: `Enable creation of shortcut routes. A received NHRP Traffic Indication will trigger the resolution and establishment of a shortcut route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
