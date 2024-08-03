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

var _ helpers.VyosTopResourceDataModel = &ProtocolsOspfAreaRange{}

// ProtocolsOspfAreaRange describes the resource data model.
type ProtocolsOspfAreaRange struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsOspfAreaRangeCost         types.Number `tfsdk:"cost" vyos:"cost,omitempty"`
	LeafProtocolsOspfAreaRangeNotAdvertise types.Bool   `tfsdk:"not_advertise" vyos:"not-advertise,omitempty"`
	LeafProtocolsOspfAreaRangeSubstitute   types.String `tfsdk:"substitute" vyos:"substitute,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsOspfAreaRange) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsOspfAreaRange) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsOspfAreaRange) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfAreaRange) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"range",
		o.SelfIdentifier.Attributes()["range"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsOspfAreaRange) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"ospf",

		"area",

		o.SelfIdentifier.Attributes()["area"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsOspfAreaRange) GetVyosNamedParentPath() []string {
	return []string{
		"protocols",

		"ospf",

		"area",

		o.SelfIdentifier.Attributes()["area"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAreaRange) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"range": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `Summarize routes matching a prefix (border routers only)

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  Area range prefix  |
`,
						Description: `Summarize routes matching a prefix (border routers only)

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  Area range prefix  |
`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in range, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  range, value must match: ^[a-zA-Z0-9-_]*$",
								),
							),
						},
					},

					"area": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `OSPF area settings

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  u32     |  OSPF area number in decimal notation         |
    |  ipv4    |  OSPF area number in dotted decimal notation  |
`,
						Description: `OSPF area settings

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  u32     |  OSPF area number in decimal notation         |
    |  ipv4    |  OSPF area number in dotted decimal notation  |
`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in area, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  area, value must match: ^[a-zA-Z0-9-_]*$",
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

		"cost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric for this range

    |  Format      |  Description            |
    |--------------|-------------------------|
    |  0-16777215  |  Metric for this range  |
`,
			Description: `Metric for this range

    |  Format      |  Description            |
    |--------------|-------------------------|
    |  0-16777215  |  Metric for this range  |
`,
		},

		"not_advertise": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not advertise this range

`,
			Description: `Do not advertise this range

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"substitute": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise area range as another prefix

    |  Format   |  Description                             |
    |-----------|------------------------------------------|
    |  ipv4net  |  Advertise area range as another prefix  |
`,
			Description: `Advertise area range as another prefix

    |  Format   |  Description                             |
    |-----------|------------------------------------------|
    |  ipv4net  |  Advertise area range as another prefix  |
`,
		},

		// Nodes

	}
}
