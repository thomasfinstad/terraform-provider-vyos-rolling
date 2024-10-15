// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceDhcpvsixServerSharedNetworkNameSubnetRange{}

// ServiceDhcpvsixServerSharedNetworkNameSubnetRange describes the resource data model.
type ServiceDhcpvsixServerSharedNetworkNameSubnetRange struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDhcpvsixServerSharedNetworkNameSubnetRangePrefix types.String `tfsdk:"prefix" vyos:"prefix,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameSubnetRangeStart  types.String `tfsdk:"start" vyos:"start,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameSubnetRangeStop   types.String `tfsdk:"stop" vyos:"stop,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeServiceDhcpvsixServerSharedNetworkNameSubnetRangeOption *ServiceDhcpvsixServerSharedNetworkNameSubnetRangeOption `tfsdk:"option" vyos:"option,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetRange) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetRange) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetRange) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetRange) GetVyosPath() []string {
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
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetRange) GetVyosParentPath() []string {
	return []string{
		"service",

		"dhcpv6-server",

		"shared-network-name",

		o.SelfIdentifier.Attributes()["shared_network_name"].(types.String).ValueString(),

		"subnet",

		o.SelfIdentifier.Attributes()["subnet"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetRange) GetVyosNamedParentPath() []string {
	return []string{
		"service",

		"dhcpv6-server",

		"shared-network-name",

		o.SelfIdentifier.Attributes()["shared_network_name"].(types.String).ValueString(),

		"subnet",

		o.SelfIdentifier.Attributes()["subnet"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpvsixServerSharedNetworkNameSubnetRange) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"range": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Parameters setting ranges for assigning IPv6 addresses

`,
					Description: `Parameters setting ranges for assigning IPv6 addresses

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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  range, value must match: ^[.:a-zA-Z0-9-_]+$",
							),
						),
					},
				},

				"shared_network_name": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `DHCPv6 shared network name

`,
					Description: `DHCPv6 shared network name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in shared_network_name, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  shared_network_name, value must match: ^[.:a-zA-Z0-9-_]+$",
							),
						),
					},
				},

				"subnet": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `IPv6 DHCP subnet for this shared network

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
					Description: `IPv6 DHCP subnet for this shared network

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in subnet, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  subnet, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 prefix defining range of addresses to assign

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `IPv6 prefix defining range of addresses to assign

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
		},

		"start": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `First in range of consecutive IPv6 addresses to assign

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv6    |  IPv6 address  |
`,
			Description: `First in range of consecutive IPv6 addresses to assign

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv6    |  IPv6 address  |
`,
		},

		"stop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Last in range of consecutive IPv6 addresses

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv6    |  IPv6 address  |
`,
			Description: `Last in range of consecutive IPv6 addresses

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv6    |  IPv6 address  |
`,
		},

		// Nodes

		"option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpvsixServerSharedNetworkNameSubnetRangeOption{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCPv6 option

`,
			Description: `DHCPv6 option

`,
		},
	}
}
