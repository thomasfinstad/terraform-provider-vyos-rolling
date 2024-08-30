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

var _ helpers.VyosTopResourceDataModel = &ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress{}

// ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddressAsSet       types.Bool   `tfsdk:"as_set" vyos:"as-set,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddressRouteMap    types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddressSummaryOnly types.Bool   `tfsdk:"summary_only" vyos:"summary-only,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"aggregate-address",
		o.SelfIdentifier.Attributes()["aggregate_address"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"bgp",

		"address-family",

		"ipv6-unicast",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"aggregate_address": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `BGP aggregate network

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv6net  |  Aggregate network  |
`,
					Description: `BGP aggregate network

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv6net  |  Aggregate network  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in aggregate_address, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  aggregate_address, value must match: ^[a-zA-Z0-9-_]*$",
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

		"as_set": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Generate AS-set path information for this aggregate address

`,
			Description: `Generate AS-set path information for this aggregate address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Specify route-map name to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		"summary_only": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Announce the aggregate summary network only

`,
			Description: `Announce the aggregate summary network only

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
