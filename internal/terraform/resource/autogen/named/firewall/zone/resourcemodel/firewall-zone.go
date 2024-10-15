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

var _ helpers.VyosTopResourceDataModel = &FirewallZone{}

// FirewallZone describes the resource data model.
type FirewallZone struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafFirewallZoneDescrIPtion   types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallZoneDefaultLog    types.Bool   `tfsdk:"default_log" vyos:"default-log,omitempty"`
	LeafFirewallZoneDefaultAction types.String `tfsdk:"default_action" vyos:"default-action,omitempty"`
	LeafFirewallZoneInterface     types.List   `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafFirewallZoneLocalZone     types.Bool   `tfsdk:"local_zone" vyos:"local-zone,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagFirewallZoneFrom bool `tfsdk:"-" vyos:"from,child"`

	// Nodes
	NodeFirewallZoneIntraZoneFiltering *FirewallZoneIntraZoneFiltering `tfsdk:"intra_zone_filtering" vyos:"intra-zone-filtering,omitempty"`
}

// SetID configures the resource ID
func (o *FirewallZone) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallZone) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallZone) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallZone) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"zone",
		o.SelfIdentifier.Attributes()["zone"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallZone) GetVyosParentPath() []string {
	return []string{
		"firewall",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallZone) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZone) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"zone": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Zone-policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Zone name    |
`,
					Description: `Zone-policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Zone name    |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in zone, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  zone, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"default_log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting default-action

`,
			Description: `Log packets hitting default-action

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"default_action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default-action for traffic coming into this zone

    |  Format  |  Description             |
    |----------|--------------------------|
    |  drop    |  Drop silently           |
    |  reject  |  Drop and notify source  |
`,
			Description: `Default-action for traffic coming into this zone

    |  Format  |  Description             |
    |----------|--------------------------|
    |  drop    |  Drop silently           |
    |  reject  |  Drop and notify source  |
`,

			// Default:          stringdefault.StaticString(`drop`),
			Computed: true,
		},

		"interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface associated with zone

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  txt     |  Interface associated with zone  |
    |  vrf     |  VRF associated with zone        |
`,
			Description: `Interface associated with zone

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  txt     |  Interface associated with zone  |
    |  vrf     |  VRF associated with zone        |
`,
		},

		"local_zone": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Zone to be local-zone

`,
			Description: `Zone to be local-zone

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"intra_zone_filtering": schema.SingleNestedAttribute{
			Attributes: FirewallZoneIntraZoneFiltering{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Intra-zone filtering

`,
			Description: `Intra-zone filtering

`,
		},
	}
}
