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

var _ helpers.VyosTopResourceDataModel = &ServiceEventHandlerEventScrIPtEnvironment{}

// ServiceEventHandlerEventScrIPtEnvironment describes the resource data model.
type ServiceEventHandlerEventScrIPtEnvironment struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceEventHandlerEventScrIPtEnvironmentValue types.String `tfsdk:"value" vyos:"value,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceEventHandlerEventScrIPtEnvironment) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceEventHandlerEventScrIPtEnvironment) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceEventHandlerEventScrIPtEnvironment) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceEventHandlerEventScrIPtEnvironment) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"environment",
		o.SelfIdentifier.Attributes()["environment"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceEventHandlerEventScrIPtEnvironment) GetVyosParentPath() []string {
	return []string{
		"service",

		"event-handler",

		"event",

		o.SelfIdentifier.Attributes()["event"].(types.String).ValueString(),

		"script",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceEventHandlerEventScrIPtEnvironment) GetVyosNamedParentPath() []string {
	return []string{
		"service",

		"event-handler",

		"event",

		o.SelfIdentifier.Attributes()["event"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceEventHandlerEventScrIPtEnvironment) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"environment": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Script environment arguments

`,
					Description: `Script environment arguments

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in environment, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  environment, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"event": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Event handler name

`,
					Description: `Event handler name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in event, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  event, value must match: ^[a-zA-Z0-9-_]*$",
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

		"value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Environment value

`,
			Description: `Environment value

`,
		},

		// Nodes

	}
}
