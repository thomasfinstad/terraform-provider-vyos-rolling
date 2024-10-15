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

var _ helpers.VyosTopResourceDataModel = &ProtocolsStaticTableRouteInterface{}

// ProtocolsStaticTableRouteInterface describes the resource data model.
type ProtocolsStaticTableRouteInterface struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsStaticTableRouteInterfaceDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafProtocolsStaticTableRouteInterfaceDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafProtocolsStaticTableRouteInterfaceVrf      types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsStaticTableRouteInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsStaticTableRouteInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsStaticTableRouteInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticTableRouteInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsStaticTableRouteInterface) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"static",

		"table",

		o.SelfIdentifier.Attributes()["table"].(types.Number).ValueBigFloat().String(),

		"route",

		o.SelfIdentifier.Attributes()["route"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsStaticTableRouteInterface) GetVyosNamedParentPath() []string {
	return []string{
		"protocols",

		"static",

		"table",

		o.SelfIdentifier.Attributes()["table"].(types.Number).ValueBigFloat().String(),

		"route",

		o.SelfIdentifier.Attributes()["route"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticTableRouteInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"interface": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Next-hop IPv4 router interface

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Gateway interface name  |
`,
					Description: `Next-hop IPv4 router interface

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Gateway interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in interface, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  interface, value must match: ^[.:a-zA-Z0-9-_]+$",
							),
						),
					},
				},

				"table": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Policy route table number

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-200   |  Policy route table number  |
`,
					Description: `Policy route table number

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-200   |  Policy route table number  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				"route": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Static IPv4 route

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  IPv4 static route  |
`,
					Description: `Static IPv4 route

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  IPv4 static route  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in route, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  route, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for this route

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-255   |  Distance for this route  |
`,
			Description: `Distance for this route

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-255   |  Distance for this route  |
`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF to leak route

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Name of VRF to leak to  |
`,
			Description: `VRF to leak route

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Name of VRF to leak to  |
`,
		},

		// Nodes

	}
}
