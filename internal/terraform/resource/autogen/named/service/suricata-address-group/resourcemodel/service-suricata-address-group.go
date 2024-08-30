// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &ServiceSURIcataAddressGroup{}

// ServiceSURIcataAddressGroup describes the resource data model.
type ServiceSURIcataAddressGroup struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceSURIcataAddressGroupAddress types.List `tfsdk:"address" vyos:"address,omitempty"`
	LeafServiceSURIcataAddressGroupGroup   types.List `tfsdk:"group" vyos:"group,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceSURIcataAddressGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceSURIcataAddressGroup) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceSURIcataAddressGroup) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSURIcataAddressGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"address-group",
		o.SelfIdentifier.Attributes()["address_group"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceSURIcataAddressGroup) GetVyosParentPath() []string {
	return []string{
		"service",

		"suricata",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceSURIcataAddressGroup) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSURIcataAddressGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"address_group": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Address group name

`,
					Description: `Address group name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in address_group, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  address_group, value must match: ^[a-zA-Z0-9-_]*$",
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

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address or subnet

    |  Format    |  Description                                      |
    |------------|---------------------------------------------------|
    |  ipv4      |  IPv4 address to match                            |
    |  ipv6      |  IPv6 address to match                            |
    |  ipv4net   |  IPv4 prefix to match                             |
    |  ipv6net   |  IPv6 prefix to match                             |
    |  !ipv4     |  Exclude the specified IPv4 address from matches  |
    |  !ipv6     |  Exclude the specified IPv6 address from matches  |
    |  !ipv4net  |  Exclude the specified IPv6 prefix from matches   |
    |  !ipv6net  |  Exclude the specified IPv6 prefix from matches   |
`,
			Description: `IP address or subnet

    |  Format    |  Description                                      |
    |------------|---------------------------------------------------|
    |  ipv4      |  IPv4 address to match                            |
    |  ipv6      |  IPv6 address to match                            |
    |  ipv4net   |  IPv4 prefix to match                             |
    |  ipv6net   |  IPv6 prefix to match                             |
    |  !ipv4     |  Exclude the specified IPv4 address from matches  |
    |  !ipv6     |  Exclude the specified IPv6 address from matches  |
    |  !ipv4net  |  Exclude the specified IPv6 prefix from matches   |
    |  !ipv6net  |  Exclude the specified IPv6 prefix from matches   |
`,
		},

		"group": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Address group

    |  Format  |  Description                                       |
    |----------|----------------------------------------------------|
    |  txt     |  Address group to match                            |
    |  !txt    |  Exclude the specified address group from matches  |
`,
			Description: `Address group

    |  Format  |  Description                                       |
    |----------|----------------------------------------------------|
    |  txt     |  Address group to match                            |
    |  !txt    |  Exclude the specified address group from matches  |
`,
		},

		// Nodes

	}
}
