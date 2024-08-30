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

var _ helpers.VyosTopResourceDataModel = &SystemFlowAccountingNetflowServer{}

// SystemFlowAccountingNetflowServer describes the resource data model.
type SystemFlowAccountingNetflowServer struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemFlowAccountingNetflowServerPort types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *SystemFlowAccountingNetflowServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemFlowAccountingNetflowServer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemFlowAccountingNetflowServer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemFlowAccountingNetflowServer) GetVyosPath() []string {
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
func (o *SystemFlowAccountingNetflowServer) GetVyosParentPath() []string {
	return []string{
		"system",

		"flow-accounting",

		"netflow",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *SystemFlowAccountingNetflowServer) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemFlowAccountingNetflowServer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `NetFlow destination server

    |  Format  |  Description                    |
    |----------|---------------------------------|
    |  ipv4    |  IPv4 server to export NetFlow  |
    |  ipv6    |  IPv6 server to export NetFlow  |
`,
					Description: `NetFlow destination server

    |  Format  |  Description                    |
    |----------|---------------------------------|
    |  ipv4    |  IPv4 server to export NetFlow  |
    |  ipv6    |  IPv6 server to export NetFlow  |
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

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `NetFlow port number

    |  Format      |  Description          |
    |--------------|-----------------------|
    |  1025-65535  |  NetFlow port number  |
`,
			Description: `NetFlow port number

    |  Format      |  Description          |
    |--------------|-----------------------|
    |  1025-65535  |  NetFlow port number  |
`,

			// Default:          stringdefault.StaticString(`2055`),
			Computed: true,
		},

		// Nodes

	}
}
