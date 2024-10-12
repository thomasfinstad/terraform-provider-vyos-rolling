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

var _ helpers.VyosTopResourceDataModel = &ProtocolsStaticRoutesixNextHop{}

// ProtocolsStaticRoutesixNextHop describes the resource data model.
type ProtocolsStaticRoutesixNextHop struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsStaticRoutesixNextHopDisable   types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafProtocolsStaticRoutesixNextHopDistance  types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafProtocolsStaticRoutesixNextHopInterface types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafProtocolsStaticRoutesixNextHopSegments  types.String `tfsdk:"segments" vyos:"segments,omitempty"`
	LeafProtocolsStaticRoutesixNextHopVrf       types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsStaticRoutesixNextHopBfd *ProtocolsStaticRoutesixNextHopBfd `tfsdk:"bfd" vyos:"bfd,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsStaticRoutesixNextHop) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsStaticRoutesixNextHop) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsStaticRoutesixNextHop) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticRoutesixNextHop) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"next-hop",
		o.SelfIdentifier.Attributes()["next_hop"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsStaticRoutesixNextHop) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"static",

		"route6",

		o.SelfIdentifier.Attributes()["route6"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsStaticRoutesixNextHop) GetVyosNamedParentPath() []string {
	return []string{
		"protocols",

		"static",

		"route6",

		o.SelfIdentifier.Attributes()["route6"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticRoutesixNextHop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"next_hop": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `IPv6 gateway address

    |  Format  |  Description           |
    |----------|------------------------|
    |  ipv6    |  Next-hop IPv6 router  |
`,
					Description: `IPv6 gateway address

    |  Format  |  Description           |
    |----------|------------------------|
    |  ipv6    |  Next-hop IPv6 router  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in next_hop, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  next_hop, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"route6": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Static IPv6 route

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv6net  |  IPv6 static route  |
`,
					Description: `Static IPv6 route

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv6net  |  IPv6 static route  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in route6, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  route6, value must match: ^[a-zA-Z0-9-_]*$",
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

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway interface name

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Gateway interface name  |
`,
			Description: `Gateway interface name

    |  Format  |  Description             |
    |----------|--------------------------|
    |  txt     |  Gateway interface name  |
`,
		},

		"segments": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `SRv6 segments

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Segs (SIDs)  |
`,
			Description: `SRv6 segments

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Segs (SIDs)  |
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

		"bfd": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticRoutesixNextHopBfd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BFD monitoring

`,
			Description: `BFD monitoring

`,
		},
	}
}
