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

var _ helpers.VyosTopResourceDataModel = &ServiceDhcpServerSharedNetworkNameSubnetRange{}

// ServiceDhcpServerSharedNetworkNameSubnetRange describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnetRange struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetRangeStart types.String `tfsdk:"start" vyos:"start,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetRangeStop  types.String `tfsdk:"stop" vyos:"stop,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeServiceDhcpServerSharedNetworkNameSubnetRangeOption *ServiceDhcpServerSharedNetworkNameSubnetRangeOption `tfsdk:"option" vyos:"option,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceDhcpServerSharedNetworkNameSubnetRange) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDhcpServerSharedNetworkNameSubnetRange) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDhcpServerSharedNetworkNameSubnetRange) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpServerSharedNetworkNameSubnetRange) GetVyosPath() []string {
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
func (o *ServiceDhcpServerSharedNetworkNameSubnetRange) GetVyosParentPath() []string {
	return []string{
		"service",

		"dhcp-server",

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
func (o *ServiceDhcpServerSharedNetworkNameSubnetRange) GetVyosNamedParentPath() []string {
	return []string{
		"service",

		"dhcp-server",

		"shared-network-name",

		o.SelfIdentifier.Attributes()["shared_network_name"].(types.String).ValueString(),

		"subnet",

		o.SelfIdentifier.Attributes()["subnet"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetRange) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `DHCP lease range

`,
					Description: `DHCP lease range

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
					MarkdownDescription: `Name of DHCP shared network

`,
					Description: `Name of DHCP shared network

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
					MarkdownDescription: `DHCP subnet for shared network

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
`,
					Description: `DHCP subnet for shared network

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
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

		"start": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `First IP address for DHCP lease range

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  IPv4 start address of pool  |
`,
			Description: `First IP address for DHCP lease range

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  IPv4 start address of pool  |
`,
		},

		"stop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Last IP address for DHCP lease range

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  IPv4 end address of pool  |
`,
			Description: `Last IP address for DHCP lease range

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  IPv4 end address of pool  |
`,
		},

		// Nodes

		"option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpServerSharedNetworkNameSubnetRangeOption{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCP option

`,
			Description: `DHCP option

`,
		},
	}
}
