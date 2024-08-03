// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource{}

// ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource describes the resource data model.
type ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsStaticTableRoutesixNextHopBfdMultiHopSourceProfile types.String `tfsdk:"profile" vyos:"profile,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"source",
		o.SelfIdentifier.Attributes()["source"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"static",

		"table",

		o.SelfIdentifier.Attributes()["table"].(types.Number).ValueBigFloat().String(),

		"route6",

		o.SelfIdentifier.Attributes()["route6"].(types.String).ValueString(),

		"next-hop",

		o.SelfIdentifier.Attributes()["next_hop"].(types.String).ValueString(),

		"bfd",

		"multi-hop",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource) GetVyosNamedParentPath() []string {
	return []string{
		"protocols",

		"static",

		"table",

		o.SelfIdentifier.Attributes()["table"].(types.Number).ValueBigFloat().String(),

		"route6",

		o.SelfIdentifier.Attributes()["route6"].(types.String).ValueString(),

		"next-hop",

		o.SelfIdentifier.Attributes()["next_hop"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticTableRoutesixNextHopBfdMultiHopSource) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"source": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `Use source for BFD session

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
    |  ipv6    |  IPv6 source address  |
`,
						Description: `Use source for BFD session

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
    |  ipv6    |  IPv6 source address  |
`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in source, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  source, value must match: ^[a-zA-Z0-9-_]*$",
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
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"profile": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
			Description: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
		},

		// Nodes

	}
}
