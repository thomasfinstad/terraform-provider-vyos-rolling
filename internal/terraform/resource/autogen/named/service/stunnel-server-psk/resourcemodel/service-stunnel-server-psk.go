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

var _ helpers.VyosTopResourceDataModel = &ServiceStunnelServerPsk{}

// ServiceStunnelServerPsk describes the resource data model.
type ServiceStunnelServerPsk struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceStunnelServerPskID     types.String `tfsdk:"id_param" vyos:"id,omitempty"`
	LeafServiceStunnelServerPskSecret types.String `tfsdk:"secret" vyos:"secret,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceStunnelServerPsk) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceStunnelServerPsk) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceStunnelServerPsk) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceStunnelServerPsk) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"psk",
		o.SelfIdentifier.Attributes()["psk"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceStunnelServerPsk) GetVyosParentPath() []string {
	return []string{
		"service",

		"stunnel",

		"server",

		o.SelfIdentifier.Attributes()["server"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceStunnelServerPsk) GetVyosNamedParentPath() []string {
	return []string{
		"service",

		"stunnel",

		"server",

		o.SelfIdentifier.Attributes()["server"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceStunnelServerPsk) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"psk": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Pre-shared key name

`,
					Description: `Pre-shared key name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in psk, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  psk, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"server": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Stunnel server config

`,
					Description: `Stunnel server config

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

		"id_param": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ID for authentication

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  ID used for authentication  |
`,
			Description: `ID for authentication

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  ID used for authentication  |
`,
		},

		"secret": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `pre-shared secret key

    |  Format  |  Description                                                                                                                |
    |----------|-----------------------------------------------------------------------------------------------------------------------------|
    |  txt     |  pre-shared secret key are required to be at least 16 bytes long, which implies at least 32 characters for hexadecimal key  |
`,
			Description: `pre-shared secret key

    |  Format  |  Description                                                                                                                |
    |----------|-----------------------------------------------------------------------------------------------------------------------------|
    |  txt     |  pre-shared secret key are required to be at least 16 bytes long, which implies at least 32 characters for hexadecimal key  |
`,
		},

		// Nodes

	}
}
