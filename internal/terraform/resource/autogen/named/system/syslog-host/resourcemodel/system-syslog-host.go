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

var _ helpers.VyosTopResourceDataModel = &SystemSyslogHost{}

// SystemSyslogHost describes the resource data model.
type SystemSyslogHost struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemSyslogHostPort     types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafSystemSyslogHostProtocol types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagSystemSyslogHostFacility bool `tfsdk:"-" vyos:"facility,child"`

	// Nodes
	NodeSystemSyslogHostFormat *SystemSyslogHostFormat `tfsdk:"format" vyos:"format,omitempty"`
}

// SetID configures the resource ID
func (o *SystemSyslogHost) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemSyslogHost) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemSyslogHost) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSyslogHost) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"host",
		o.SelfIdentifier.Attributes()["host"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemSyslogHost) GetVyosParentPath() []string {
	return []string{
		"system",

		"syslog",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *SystemSyslogHost) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogHost) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `Logging to remote host

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  ipv4      |  Remote syslog server IPv4 address  |
    |  ipv6      |  Remote syslog server IPv6 address  |
    |  hostname  |  Remote syslog server FQDN          |
`,
						Description: `Logging to remote host

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  ipv4      |  Remote syslog server IPv4 address  |
    |  ipv6      |  Remote syslog server IPv6 address  |
    |  hostname  |  Remote syslog server FQDN          |
`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in host, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  host, value must match: ^[a-zA-Z0-9-_]*$",
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

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,

			// Default:          stringdefault.StaticString(`514`),
			Computed: true,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to be used (TCP/UDP)

    |  Format  |  Description          |
    |----------|-----------------------|
    |  udp     |  Listen protocol UDP  |
    |  tcp     |  Listen protocol TCP  |
`,
			Description: `Protocol to be used (TCP/UDP)

    |  Format  |  Description          |
    |----------|-----------------------|
    |  udp     |  Listen protocol UDP  |
    |  tcp     |  Listen protocol TCP  |
`,

			// Default:          stringdefault.StaticString(`udp`),
			Computed: true,
		},

		// Nodes

		"format": schema.SingleNestedAttribute{
			Attributes: SystemSyslogHostFormat{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Logging format

`,
			Description: `Logging format

`,
		},
	}
}