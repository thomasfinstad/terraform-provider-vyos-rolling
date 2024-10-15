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

var _ helpers.VyosTopResourceDataModel = &SystemLoginRadiusServer{}

// SystemLoginRadiusServer describes the resource data model.
type SystemLoginRadiusServer struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemLoginRadiusServerDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafSystemLoginRadiusServerKey      types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafSystemLoginRadiusServerPort     types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafSystemLoginRadiusServerTimeout  types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafSystemLoginRadiusServerPriority types.Number `tfsdk:"priority" vyos:"priority,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *SystemLoginRadiusServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemLoginRadiusServer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemLoginRadiusServer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemLoginRadiusServer) GetVyosPath() []string {
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
func (o *SystemLoginRadiusServer) GetVyosParentPath() []string {
	return []string{
		"system",

		"login",

		"radius",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *SystemLoginRadiusServer) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginRadiusServer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `RADIUS server configuration

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  RADIUS server IPv4 address  |
    |  ipv6    |  RADIUS server IPv6 address  |
`,
					Description: `RADIUS server configuration

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  RADIUS server IPv4 address  |
    |  ipv6    |  RADIUS server IPv6 address  |
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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  server, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shared secret key

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  Password string (key)  |
`,
			Description: `Shared secret key

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  Password string (key)  |
`,
		},

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

			// Default:          stringdefault.StaticString(`1812`),
			Computed: true,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session timeout

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  1-240   |  Session timeout in seconds (default: 2)  |
`,
			Description: `Session timeout

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  1-240   |  Session timeout in seconds (default: 2)  |
`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Server priority

    |  Format  |  Description      |
    |----------|-------------------|
    |  1-255   |  Server priority  |
`,
			Description: `Server priority

    |  Format  |  Description      |
    |----------|-------------------|
    |  1-255   |  Server priority  |
`,

			// Default:          stringdefault.StaticString(`255`),
			Computed: true,
		},

		// Nodes

	}
}
