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

var _ helpers.VyosTopResourceDataModel = &ServiceNtpServer{}

// ServiceNtpServer describes the resource data model.
type ServiceNtpServer struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceNtpServerNoselect types.Bool `tfsdk:"noselect" vyos:"noselect,omitempty"`
	LeafServiceNtpServerNts      types.Bool `tfsdk:"nts" vyos:"nts,omitempty"`
	LeafServiceNtpServerPool     types.Bool `tfsdk:"pool" vyos:"pool,omitempty"`
	LeafServiceNtpServerPrefer   types.Bool `tfsdk:"prefer" vyos:"prefer,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceNtpServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceNtpServer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceNtpServer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceNtpServer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"server",
		o.SelfIdentifier.Attributes()["server"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceNtpServer) GetVyosParentPath() []string {
	return []string{
		"service",

		"ntp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceNtpServer) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceNtpServer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"server": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Network Time Protocol (NTP) server

    |  Format    |  Description                                |
    |------------|---------------------------------------------|
    |  ipv4      |  IP address of NTP server                   |
    |  ipv6      |  IPv6 address of NTP server                 |
    |  hostname  |  Fully qualified domain name of NTP server  |
`,
					Description: `Network Time Protocol (NTP) server

    |  Format    |  Description                                |
    |------------|---------------------------------------------|
    |  ipv4      |  IP address of NTP server                   |
    |  ipv6      |  IPv6 address of NTP server                 |
    |  hostname  |  Fully qualified domain name of NTP server  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in server, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  server, value must match: ^[a-zA-Z0-9-_]*$",
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

		"noselect": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Marks the server as unused

`,
			Description: `Marks the server as unused

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"nts": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Network Time Security (NTS) for the server

`,
			Description: `Enable Network Time Security (NTS) for the server

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"pool": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Associate with a number of remote servers

`,
			Description: `Associate with a number of remote servers

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"prefer": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Marks the server as preferred

`,
			Description: `Marks the server as preferred

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
