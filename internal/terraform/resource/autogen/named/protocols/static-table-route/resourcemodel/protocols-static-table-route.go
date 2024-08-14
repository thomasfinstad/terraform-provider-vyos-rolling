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

var _ helpers.VyosTopResourceDataModel = &ProtocolsStaticTableRoute{}

// ProtocolsStaticTableRoute describes the resource data model.
type ProtocolsStaticTableRoute struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsStaticTableRouteDhcpInterface types.String `tfsdk:"dhcp_interface" vyos:"dhcp-interface,omitempty"`
	LeafProtocolsStaticTableRouteDescrIPtion   types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagProtocolsStaticTableRouteInterface bool `tfsdk:"-" vyos:"interface,child"`

	ExistsTagProtocolsStaticTableRouteNextHop bool `tfsdk:"-" vyos:"next-hop,child"`

	// Nodes
	NodeProtocolsStaticTableRouteBlackhole *ProtocolsStaticTableRouteBlackhole `tfsdk:"blackhole" vyos:"blackhole,omitempty"`
	NodeProtocolsStaticTableRouteReject    *ProtocolsStaticTableRouteReject    `tfsdk:"reject" vyos:"reject,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsStaticTableRoute) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsStaticTableRoute) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsStaticTableRoute) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticTableRoute) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"route",
		o.SelfIdentifier.Attributes()["route"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsStaticTableRoute) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"static",

		"table",

		o.SelfIdentifier.Attributes()["table"].(types.Number).ValueBigFloat().String(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsStaticTableRoute) GetVyosNamedParentPath() []string {
	return []string{
		"protocols",

		"static",

		"table",

		o.SelfIdentifier.Attributes()["table"].(types.Number).ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticTableRoute) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
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
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  route, value must match: ^[a-zA-Z0-9-_]*$",
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
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"dhcp_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP interface supplying next-hop IP address

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  DHCP interface name  |
`,
			Description: `DHCP interface supplying next-hop IP address

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  DHCP interface name  |
`,
		},

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

		// Nodes

		"blackhole": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticTableRouteBlackhole{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Silently discard pkts when matched

`,
			Description: `Silently discard pkts when matched

`,
		},

		"reject": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticTableRouteReject{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Emit an ICMP unreachable when matched

`,
			Description: `Emit an ICMP unreachable when matched

`,
		},
	}
}