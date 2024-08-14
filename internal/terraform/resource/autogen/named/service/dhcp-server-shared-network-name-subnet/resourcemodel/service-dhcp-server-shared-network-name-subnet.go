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

var _ helpers.VyosTopResourceDataModel = &ServiceDhcpServerSharedNetworkNameSubnet{}

// ServiceDhcpServerSharedNetworkNameSubnet describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnet struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion    types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetDisable        types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetExclude        types.List   `tfsdk:"exclude" vyos:"exclude,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetIgnoreClientID types.Bool   `tfsdk:"ignore_client_id" vyos:"ignore-client-id,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetLease          types.Number `tfsdk:"lease" vyos:"lease,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetSubnetID       types.Number `tfsdk:"subnet_id" vyos:"subnet-id,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagServiceDhcpServerSharedNetworkNameSubnetRange bool `tfsdk:"-" vyos:"range,child"`

	ExistsTagServiceDhcpServerSharedNetworkNameSubnetStaticMapping bool `tfsdk:"-" vyos:"static-mapping,child"`

	// Nodes
	NodeServiceDhcpServerSharedNetworkNameSubnetOption *ServiceDhcpServerSharedNetworkNameSubnetOption `tfsdk:"option" vyos:"option,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceDhcpServerSharedNetworkNameSubnet) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDhcpServerSharedNetworkNameSubnet) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDhcpServerSharedNetworkNameSubnet) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpServerSharedNetworkNameSubnet) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"subnet",
		o.SelfIdentifier.Attributes()["subnet"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDhcpServerSharedNetworkNameSubnet) GetVyosParentPath() []string {
	return []string{
		"service",

		"dhcp-server",

		"shared-network-name",

		o.SelfIdentifier.Attributes()["shared_network_name"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDhcpServerSharedNetworkNameSubnet) GetVyosNamedParentPath() []string {
	return []string{
		"service",

		"dhcp-server",

		"shared-network-name",

		o.SelfIdentifier.Attributes()["shared_network_name"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnet) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
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
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  subnet, value must match: ^[a-zA-Z0-9-_]*$",
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
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  shared_network_name, value must match: ^[a-zA-Z0-9-_]*$",
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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"exclude": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address to exclude from DHCP lease range

    |  Format  |  Description                               |
    |----------|--------------------------------------------|
    |  ipv4    |  IPv4 address to exclude from lease range  |
`,
			Description: `IP address to exclude from DHCP lease range

    |  Format  |  Description                               |
    |----------|--------------------------------------------|
    |  ipv4    |  IPv4 address to exclude from lease range  |
`,
		},

		"ignore_client_id": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore client identifier for lease lookups

`,
			Description: `Ignore client identifier for lease lookups

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"lease": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Lease timeout in seconds

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  u32     |  DHCP lease time in seconds  |
`,
			Description: `Lease timeout in seconds

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  u32     |  DHCP lease time in seconds  |
`,

			// Default:          stringdefault.StaticString(`86400`),
			Computed: true,
		},

		"subnet_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Unique ID mapped to leases in the lease file

    |  Format  |  Description       |
    |----------|--------------------|
    |  u32     |  Unique subnet ID  |
`,
			Description: `Unique ID mapped to leases in the lease file

    |  Format  |  Description       |
    |----------|--------------------|
    |  u32     |  Unique subnet ID  |
`,
		},

		// Nodes

		"option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpServerSharedNetworkNameSubnetOption{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCP option

`,
			Description: `DHCP option

`,
		},
	}
}